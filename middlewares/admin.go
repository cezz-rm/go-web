package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-web/models"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityID != 2 {
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "must admin",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
