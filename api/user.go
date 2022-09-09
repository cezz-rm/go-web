package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "get user list succeed",
	})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "login succeed",
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "register succeed",
	})
}
