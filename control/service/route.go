package service

import (
	"fmt"
	"log"
	"monitor-control/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	serPort := config.GetInt32("service.port")
	// 加载路由
	route := loadRoute()
	log.Println("Api server startup success, port is", serPort)
	route.Run(fmt.Sprintf(":%v", serPort))
}

func loadRoute() *gin.Engine {
	route := gin.Default()
	//定义默认路由
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "这是主控程序的首页",
		})
	})
	return route
}
