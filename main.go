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
	route.GET("/m1", controllers.GetM1)
	route.GET("/m2", controllers.GetM2)
	route.GET("/m3", controllers.GetM3)
	route.NoRoute(controllers.NoRoute)

	log.Println("Project start success, The program runs on port 5555")

	err := route.Run("0.0.0.0:5555")
	if err != nil {
		log.Println("Error：", err)
		return
	}
}
