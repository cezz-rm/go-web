package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSms(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "send succeed",
	})
}
