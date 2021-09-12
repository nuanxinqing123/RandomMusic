/**
 * @Author: Nuanxinqing
 * @Email：nuanxinqing@gmail.com
 * @File:  main
 * @Date: 2021/8/20 16:48
 */

package main

import (
	"RandomMusic/controllers"
	_ "RandomMusic/dataSource"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	//加载网页文件
	route.LoadHTMLGlob("./views/**/*")

	//设置session中间件
	store := cookie.NewStore([]byte("userLogin"))
	route.Use(sessions.Sessions("pt_Key", store))

	{
		//首页
		route.GET("/", controllers.GetIndex)
	}

	// 电脑端
	{
		//热歌
		route.GET("/m1", controllers.GetM1)
		//新歌
		route.GET("/m2", controllers.GetM2)
		//飙升
		route.GET("/m3", controllers.GetM3)
		//原创
		route.GET("/m4", controllers.GetM4)
		//登录页面
		route.GET("/m/login", controllers.GetMUserLogin)
		//注册页面
		route.GET("/m/register", controllers.GetMUserRegister)
		//Send验证码页面
		route.GET("/m/getpaswd", controllers.GetMPaswd)
		//改密页面
		route.GET("/m/password", controllers.GetMPassWord)
	}

	// 移动端
	{
		//热歌
		route.GET("/a1", controllers.GetA1)
		//新歌
		route.GET("/a2", controllers.GetA2)
		//飙升
		route.GET("/a3", controllers.GetA3)
		//原创
		route.GET("/a4", controllers.GetA4)
		//登录页面
		route.GET("/a/login", controllers.GetAUserLogin)
		//注册页面
		route.GET("/a/register", controllers.GetAUserRegister)
		//Send验证码页面
		route.GET("/a/getpaswd", controllers.GetAPaswd)
		//改密页面
		route.GET("/a/password", controllers.GetAPassWord)
	}

	//公共接口
	{
		//登录接口
		route.POST("/login", controllers.PostUserLogin)
		//注册接口
		route.POST("/register", controllers.PostUserRegister)
		//退出接口
		route.GET("/exit", controllers.ExitGet)
		//验证码接口
		route.POST("/emailcode", controllers.PostEmailCode)
		//改密接口
		route.POST("/getpassword", controllers.PostPaswd)
	}

	// 404错误
	route.NoRoute(controllers.NoRoute)

	log.Println("Start success, Runs on port 5555")

	err := route.Run("0.0.0.0:5555")
	if err != nil {
		log.Println("Error：", err)
		return
	}
}
