package main

import (
	"github.com/run100/go_module_test/router"
	"github.com/run100/go_module_test/utils"
)

func main() {

	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()

	r := router.Router()
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
