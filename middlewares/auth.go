package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/belovetech/e-commerce/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		user, err := ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

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

func ValidateToken(token string) (*utils.User, error) {
	claims, err := utils.VerifyJWT(token)
	if err != nil {
		return nil, err
	}
	user := &utils.User{
		ID:    claims.ID,
		Email: claims.Email,
		Role:  claims.Role,
	}

	return user, err
}
