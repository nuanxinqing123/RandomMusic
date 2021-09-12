/**
 * @Author: Nuanxinqing
 * @Email：nuanxinqing@gmail.com
 * @File:  bgControllers
 * @Date: 2021/8/20 17:36
 */

package controllers

import (
	"RandomMusic/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetIndex 首页文件
func GetIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

//获取背景信息
func getbg(num int) (url, author string) {
	//API请求Bing每日一图
	//PC尺寸：http://s.cn.bing.net/th?id=OHR.DjurdjevicaBridge_ZH-CN0284105882_1920x1080.jpg&rf=LaDigue_1920x1080.jpg&pid=hp
	//手机尺寸：https://cn.bing.com/th?id=OHR.DjurdjevicaBridge_ZH-CN0284105882_768x1366.jpg&rf=LaDigue_1920x1080.jpg&pid=hp
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
	//转换https
	bgData.Bing.Url = htps(bgData.Bing.Url)

	//判断移动端 & 电脑端
	if num == 1 {
		bgData.Bing.Url = MobileSize(bgData.Bing.Url)
	}

	return bgData.Bing.Url, bgData.Bing.Copyright
}

// MobileSize 修改移动端尺寸
func MobileSize(data string) string {
	old := "1920x1080"
	newUrl := "768x1366"
	res := strings.Replace(data, old, newUrl, -1)
	return res
}

//获取随机音乐
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

//修改http到https
func htps(data string) string {
	index := 4
	res := data[:index] + "s" + data[index:]
	return res
}

// NoRoute 404错误信息
func NoRoute(ctx *gin.Context) {
	//ctx.String(http.StatusOK, "The current page does not exist, please return to the previous page")
	ctx.HTML(http.StatusOK, "404.html", nil)
}

// XLogin 登录选择
func XLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "x_login.html", nil)
}
