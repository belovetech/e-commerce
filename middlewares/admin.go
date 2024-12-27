package middlewares

import (
	"log"
	"net/http"

	"github.com/belovetech/e-commerce/utils"
	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, exists := c.Get("currentUser")
		log.Printf("Current user: %v", currentUser)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		claims, ok := currentUser.(*utils.User)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"message": "forbidden"})
			c.Abort()
			return
		}

		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"message": "forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
