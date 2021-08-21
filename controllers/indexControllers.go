/**
 * @Author: Nuanxinqing
 * @Email：nuanxinqing@gmail.com
 * @File:  bgControllers
 * @Date: 2021/8/20 17:36
 */

package controllers

import (
	"Gin_163Music/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// GetIndex 首页文件
func GetIndex(ctx *gin.Context) {
	url, copyright := getbg()
	name, auther, picurl, mp3url := getMusic("https://api.qqsuu.cn/api/rand.music?sort=热歌榜&type=json")


	ctx.HTML(http.StatusOK,"index.html", gin.H{
		"url": url,
		"copyright": copyright,
		"name": name,
		"auther": auther,
		"picurl": picurl,
		"mp3url": mp3url,
	})
}

//获取背景信息
func getbg() (url, author string) {
	//API请求Bing每日一图
	link := "https://api.no0a.cn/api/bing/0"
	res, err := http.Get(link)
	if err != nil {
		fmt.Println("Error：", err)
		return
	}
	//解析内容
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error：", err)
		return
	}
	res.Body.Close()

	//JSON获取
	var bgData model.BgAPI
	err = json.Unmarshal(data, &bgData)
	if err != nil {
		fmt.Println("Error：", err)
		return
	}

	return bgData.Bing.Url, bgData.Bing.Copyright
}

func getMusic(url string) (name, auther, mp3url, picurl string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error：", err)
		return
	}
	//解析内容
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error：", err)
		return
	}
	res.Body.Close()

	//JSON获取
	var musicData model.MusicData
	err = json.Unmarshal(data, &musicData)
	if err != nil {
		fmt.Println("Error：", err)
		return
	}

	return musicData.Info.Name, musicData.Info.Auther, musicData.Info.Picurl, musicData.Info.Mp3url
}