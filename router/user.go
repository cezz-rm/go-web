package router

import (
	"github.com/gin-gonic/gin"

	"go-web/api"
	"go-web/middlewares"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("login", api.Login)
		UserRouter.POST("register", api.Register)
	}
}
