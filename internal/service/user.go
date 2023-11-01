package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInit(c *gin.Context) {
	data := map[string]interface{}{
		"msg": "注册成功",
	}

	c.JSON(http.StatusOK, data)
}

func UserInfo(c *gin.Context) {
	data := map[string]interface{}{
		"name": "123123",
		"age":  111,
	}

	c.JSON(http.StatusOK, data)
}
