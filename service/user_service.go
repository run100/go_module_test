package service

import (
	"github.com/gin-gonic/gin"
	"github.com/run100/go_module_test/models"
)

func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)

	data = models.GetUserList()

	c.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "用户名已注册！",
		"data":    data,
	})
}
