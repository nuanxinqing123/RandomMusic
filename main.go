/**
 * @Author: Nuanxinqing
 * @Email：nuanxinqing@gmail.com
 * @File:  main
 * @Date: 2021/8/20 16:48
 */

package main

import (
	"Gin_163Music/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	//加载网页文件
	route.LoadHTMLGlob("./views/*")

	route.GET("/", controllers.GetIndex)

	// 电脑端
	route.GET("/m1", controllers.GetM1)
	route.GET("/m2", controllers.GetM2)
	route.GET("/m3", controllers.GetM3)
	route.GET("/m4", controllers.GetM4)

	// 移动端
	route.GET("/a1", controllers.GetA1)
	route.GET("/a2", controllers.GetA2)
	route.GET("/a3", controllers.GetA3)
	route.GET("/a4", controllers.GetA4)

	// 404错误
	route.NoRoute(controllers.NoRoute)

	log.Println("Project start success, runs on port 5555")

	err := route.Run("0.0.0.0:5555")
	if err != nil {
		log.Println("Error：", err)
		return
	}
}
