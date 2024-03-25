package router

import (
	"github.com/gin-gonic/gin"
	"github.com/run100/go_module_test/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/", service.GetIndex)

	return r
}
