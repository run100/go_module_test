package router

import (
	"github.com/gin-gonic/gin"
	"github.com/run100/go_module_test/docs"
	"github.com/run100/go_module_test/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//静态资源
	r.Static("/asset", "asset/")

	r.GET("/", service.GetIndex)

	//用户模块
	r.GET("/user/getUserList", service.GetUserList)
	r.POST("/user/createUser", service.CreateUser)

	r.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	return r
}
