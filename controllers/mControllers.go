/**
 * @Author: Nuanxinqing
 * @Email：nuanxinqing@gmail.com
 * @File:  m1Controllers.go
 * @Date: 2021/8/20 19:14
 */

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetM1(ctx *gin.Context) {
	//获取session，判断用户是否登录
	isLogin := GetSession(ctx)

	url, copyright := getbg(0)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=热歌榜&type=json")

	if isLogin == true {
		userName := GetUserName(ctx)

		ctx.HTML(http.StatusOK, "pc.html", gin.H{
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
		ctx.HTML(http.StatusOK, "pc.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
		})
	}
}

func GetM2(ctx *gin.Context) {
	//获取session，判断用户是否登录
	isLogin := GetSession(ctx)

	url, copyright := getbg(0)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=新歌榜&type=json")

	if isLogin == true {
		userName := GetUserName(ctx)

		ctx.HTML(http.StatusOK, "pc.html", gin.H{
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
		ctx.HTML(http.StatusOK, "pc.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
		})
	}
}

func GetM3(ctx *gin.Context) {
	//获取session，判断用户是否登录
	isLogin := GetSession(ctx)

	url, copyright := getbg(0)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=飙升榜&type=json")

	if isLogin == true {
		userName := GetUserName(ctx)

		ctx.HTML(http.StatusOK, "pc.html", gin.H{
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
		ctx.HTML(http.StatusOK, "pc.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
		})
	}
}

func GetM4(ctx *gin.Context) {
	//获取session，判断用户是否登录
	isLogin := GetSession(ctx)

	url, copyright := getbg(0)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=原创榜&type=json")

	if isLogin == true {
		userName := GetUserName(ctx)

		ctx.HTML(http.StatusOK, "pc.html", gin.H{
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
		ctx.HTML(http.StatusOK, "pc.html", gin.H{
			"url":       url,
			"copyright": copyright,
			"name":      name,
			"auther":    auther,
			"picurl":    picurl,
			"mp3url":    mp3url,
		})
	}
}

// GetMUserLogin PC - Login
func GetMUserLogin(ctx *gin.Context) {
	url, _ := getbg(0)

	ctx.HTML(http.StatusOK, "mlogin.html", gin.H{
		"url": url,
	})
}

// GetMUserRegister PC - Register
func GetMUserRegister(ctx *gin.Context) {
	url, _ := getbg(0)

	ctx.HTML(http.StatusOK, "mregister.html", gin.H{
		"url": url,
	})
}

// GetMPaswd PC - 找回密码
func GetMPaswd(ctx *gin.Context) {
	url, _ := getbg(0)

	ctx.HTML(http.StatusOK, "mpaswd.html", gin.H{
		"url": url,
	})
}

// GetMPassWord PC - 修改密码
func GetMPassWord(ctx *gin.Context) {
	url, _ := getbg(0)

	ctx.HTML(http.StatusOK, "mpassword.html", gin.H{
		"url": url,
	})
}
