/**
 * @Author: Nuanxinqing
 * @Email：nuanxinqing@gmail.com
 * @File:  aControllers.go
 * @Date: 2021/8/31 18:44
 */

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetA1(ctx *gin.Context) {
	//获取session，判断用户是否登录
	isLogin := GetSession(ctx)
	url, copyright := getbg(1)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=热歌榜&type=json")

	if isLogin == true {
		userName := GetUserName(ctx)

		ctx.HTML(http.StatusOK, "app.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
			"IsLogin":   isLogin,
			"UserName":  userName,
		})
	} else {
		ctx.HTML(http.StatusOK, "app.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
		})
	}
}

func GetA2(ctx *gin.Context) {
	//获取session，判断用户是否登录
	isLogin := GetSession(ctx)
	url, copyright := getbg(1)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=新歌榜&type=json")

	if isLogin == true {
		userName := GetUserName(ctx)

		ctx.HTML(http.StatusOK, "app.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
			"IsLogin":   isLogin,
			"UserName":  userName,
		})
	} else {
		ctx.HTML(http.StatusOK, "app.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
		})
	}
}

func GetA3(ctx *gin.Context) {
	//获取session，判断用户是否登录
	isLogin := GetSession(ctx)
	url, copyright := getbg(1)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=飙升榜&type=json")

	if isLogin == true {
		userName := GetUserName(ctx)

		ctx.HTML(http.StatusOK, "app.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
			"IsLogin":   isLogin,
			"UserName":  userName,
		})
	} else {
		ctx.HTML(http.StatusOK, "app.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
		})
	}
}

func GetA4(ctx *gin.Context) {
	//获取session，判断用户是否登录
	isLogin := GetSession(ctx)
	url, copyright := getbg(1)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=原创榜&type=json")

	if isLogin == true {
		userName := GetUserName(ctx)

		ctx.HTML(http.StatusOK, "app.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
			"IsLogin":   isLogin,
			"UserName":  userName,
		})
	} else {
		ctx.HTML(http.StatusOK, "app.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
		})
	}
}

// GetAUserLogin 移动 - Login
func GetAUserLogin(ctx *gin.Context) {
	url, _ := getbg(1)

	ctx.HTML(http.StatusOK, "alogin.html", gin.H{
		"url": url,
	})
}

// GetAUserRegister 移动 - Register
func GetAUserRegister(ctx *gin.Context) {
	url, _ := getbg(1)

	ctx.HTML(http.StatusOK, "aregister.html", gin.H{
		"url": url,
	})
}

// GetAPaswd 移动 - 找回密码
func GetAPaswd(ctx *gin.Context) {
	url, _ := getbg(1)

	ctx.HTML(http.StatusOK, "apaswd.html", gin.H{
		"url": url,
	})
}

// GetAPassWord 移动 - 修改密码
func GetAPassWord(ctx *gin.Context) {
	url, _ := getbg(1)

	ctx.HTML(http.StatusOK, "apassword.html", gin.H{
		"url": url,
	})
}
