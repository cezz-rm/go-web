package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCaptcha(c *gin.Context) {
	// github.com/mojocn/base64Captcha
	c.JSON(http.StatusOK, gin.H{
		"captchaID": "id",
		"picPath":   "b64s",
	})
}
