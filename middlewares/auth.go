package middlewares

import (
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
		user, err := validateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

func validateToken(token string) (*utils.User, error) {
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
