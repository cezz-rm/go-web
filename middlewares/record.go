package middlewares

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ReqLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		method := c.Request.Method
		reqURL := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		get, _ := c.Get("req")
		marshal, _ := json.Marshal(get)

		log.WithFields(log.Fields{
			"status_code": statusCode,
			"client_ip":   clientIP,
			"req_method":  method,
			"req_uri":     reqURL,
			"body":        string(marshal),
		}).Info()
	}
}
