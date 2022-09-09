package initialize

import (
	"github.com/gin-gonic/gin"

	"go-web/middlewares"
	"go-web/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	APIGroup := r.Group("/api/v1")
	router.InitUserRouter(APIGroup)
	router.InitBaseRouter(APIGroup)
	return r
}
