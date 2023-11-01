package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lyico/server/internal/service"
)

func InitUserRouter(r *gin.Engine) {
	userRouter := r.Group("/user")

	userRouter.GET("/init", service.UserInit)

	userRouter.GET("/info", service.UserInfo)
}
