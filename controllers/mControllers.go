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
	url, copyright := getbg(0)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=热歌榜&type=json")

	ctx.HTML(http.StatusOK, "pc.html", gin.H{
		"url":       url,
		"copyright": copyright,
		"name":      name,
		"auther":    auther,
		"picurl":    picurl,
		"mp3url":    mp3url,
	})
}

func GetM2(ctx *gin.Context) {
	url, copyright := getbg(0)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=新歌榜&type=json")

	ctx.HTML(http.StatusOK, "pc.html", gin.H{
		"url":       url,
		"copyright": copyright,
		"name":      name,
		"auther":    auther,
		"picurl":    picurl,
		"mp3url":    mp3url,
	})
}

func GetM3(ctx *gin.Context) {
	url, copyright := getbg(0)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=飙升榜&type=json")

	ctx.HTML(http.StatusOK, "pc.html", gin.H{
		"url":       url,
		"copyright": copyright,
		"name":      name,
		"auther":    auther,
		"picurl":    picurl,
		"mp3url":    mp3url,
	})
}

func GetM4(ctx *gin.Context) {
	url, copyright := getbg(0)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=原创榜&type=json")

	ctx.HTML(http.StatusOK, "pc.html", gin.H{
		"url":       url,
		"copyright": copyright,
		"name":      name,
		"auther":    auther,
		"picurl":    picurl,
		"mp3url":    mp3url,
	})
}
