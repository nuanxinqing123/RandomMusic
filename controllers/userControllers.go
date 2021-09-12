package controllers

import (
	"RandomMusic/dataSource"
	"RandomMusic/model"
	"RandomMusic/tools"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// PostUserRegister 用户注册
func PostUserRegister(ctx *gin.Context) {
	var userRegister model.User
	err := ctx.ShouldBind(&userRegister)
	//错误处理
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": "系统出错，请稍后尝试"})
	} else {
		//判断请求是否为空
		if userRegister.Username == "" || userRegister.Password == "" || userRegister.Email == "" || tools.VerifyEmailFormat(userRegister.Email) == false {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "请勿非法请求接口",
			})
		} else {
			//判断用户名邮箱是否已注册
			dataSource.DB.Where("username = ?", userRegister.Username).First(&userRegister)
			if userRegister.Id != 0 {
				ctx.JSON(http.StatusOK, gin.H{"msg": "用户名已被使用"})
			} else {
				dataSource.DB.Where("email = ?", userRegister.Email).First(&userRegister)
				if userRegister.Id != 0 {
					ctx.JSON(http.StatusOK, gin.H{"msg": "邮箱已被使用"})
				} else {
					//Password加密
					userRegister.Password = tools.AddMD5(userRegister.Password)

					//设置注册时间
					userRegister.RegisterTime = tools.SwitchTimeStampToData(time.Now().Unix())

					//提交注册
					dataSource.DB.Create(&userRegister)

					ctx.Redirect(http.StatusMovedPermanently, "/xlogin")
				}
			}
		}
	}
}

// PostUserLogin 用户登录
func PostUserLogin(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBind(&user)
	//错误处理
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": "系统出错，请稍后尝试"})

	} else {
		//加密
		userPassword := tools.AddMD5(user.Password)
		// 查询
		dataSource.DB.Where("username = ?", user.Username).First(&user)
		//判断用户名是否存在
		if user.Id == 0 {
			ctx.JSON(http.StatusOK, gin.H{"msg": "用户名不存在"})
		} else {
			//判断密码是否正确
			if userPassword != user.Password {
				ctx.JSON(http.StatusOK, gin.H{"msg": "密码错误"})
			} else {
				/*
				   设置了session后将数据处理设置到cookie，然后再浏览器进行网络请求的时候回自动带上cookie
				   因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置
				*/
				session := sessions.Default(ctx)
				session.Set("userLogin", user.Username)
				session.Save()

				//返回登录状态
				ctx.Redirect(http.StatusMovedPermanently, "/")
			}
		}
	}
}

// ExitGet 退出登录
func ExitGet(ctx *gin.Context) {
	//清楚用户登录状态数据
	session := sessions.Default(ctx)
	session.Delete("userLogin")
	session.Clear()
	session.Save()
}

// PostEmailCode 发送邮箱验证码
func PostEmailCode(ctx *gin.Context) {
	var user model.User
	email := ctx.PostForm("email")
	user.Email = email

	//判断邮箱是否已注册
	dataSource.DB.Where("email = ?", user.Email).First(&user)
	if user.Id == 0 {
		ctx.JSON(http.StatusOK, gin.H{"msg": "当前邮箱未注册，无法找回密码"})
	} else {
		//判断是否存在验证码
		var EmailCode model.UserEmailCode
		EmailCode.Email = email
		dataSource.DB.Where("email = ?", EmailCode.Email).First(&EmailCode)
		if EmailCode.Id != 0 {
			ctx.JSON(http.StatusOK, gin.H{"msg": "存在未使用验证码，请勿刷邮件验证码"})
		} else {
			//发送邮件
			go tools.SendEmailCode(EmailCode.Email)
			ctx.Redirect(http.StatusMovedPermanently, "/m/password")
		}
	}
}

// PostPaswd POST修改密码
func PostPaswd(ctx *gin.Context) {
	//提交验证码和新密码
	code := ctx.PostForm("code")
	email := ctx.PostForm("email")
	paswd := ctx.PostForm("password")

	if code == "" || email == "" || paswd == "" {
		ctx.JSON(http.StatusOK, gin.H{"msg": "参数不完整"})
	} else {
		var EmailCode model.UserEmailCode
		EmailCode.Email = email
		//通过邮箱查询验证码过期时间
		dataSource.DB.Where("email = ?", EmailCode.Email).First(&EmailCode)
		//匹配验证码是否超出过期时间
		if int(time.Now().Unix()) > EmailCode.EndTime {
			ctx.JSON(http.StatusOK, gin.H{"msg": "验证码已过期"})
		} else {
			//验证码匹配，修改用户密码
			if EmailCode.Code != code {
				ctx.JSON(http.StatusOK, gin.H{"msg": "验证码错误"})
			} else {
				var users model.User
				//修改密码
				dataSource.DB.Where("email = ?", EmailCode.Email).First(&users)
				users.Password = tools.AddMD5(paswd)
				dataSource.DB.Model(&users).Update("password", users.Password)
				//删除验证码
				dataSource.DB.Delete(&EmailCode)

				//返回首页
				ctx.Redirect(http.StatusMovedPermanently, "/")
			}
		}
	}
}
