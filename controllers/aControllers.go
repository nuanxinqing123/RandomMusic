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
	url, copyright := getbg(1)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=热歌榜&type=json")

	ctx.HTML(http.StatusOK, "app.html", gin.H{
		"url":       url,
		"copyright": copyright,
		"name":      name,
		"auther":    auther,
		"picurl":    picurl,
		"mp3url":    mp3url,
	})
}

func GetA2(ctx *gin.Context) {
	url, copyright := getbg(1)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=新歌榜&type=json")

	ctx.HTML(http.StatusOK, "app.html", gin.H{
		"url":       url,
		"copyright": copyright,
		"name":      name,
		"auther":    auther,
		"picurl":    picurl,
		"mp3url":    mp3url,
	})
}

func GetA3(ctx *gin.Context) {
	url, copyright := getbg(1)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=飙升榜&type=json")

	ctx.HTML(http.StatusOK, "app.html", gin.H{
		"url":       url,
		"copyright": copyright,
		"name":      name,
		"auther":    auther,
		"picurl":    picurl,
		"mp3url":    mp3url,
	})
}

func GetA4(ctx *gin.Context) {
	url, copyright := getbg(1)
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=原创榜&type=json")

	ctx.HTML(http.StatusOK, "app.html", gin.H{
		"url":       url,
		"copyright": copyright,
		"name":      name,
		"auther":    auther,
		"picurl":    picurl,
		"mp3url":    mp3url,
	})
}
