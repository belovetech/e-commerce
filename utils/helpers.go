package utils

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(c *gin.Context) int {
	user, _ := c.Get("currentUser")
	return user.(*User).ID
}

func NormalizeStatus(status string) string {
	switch strings.ToLower(status) {
	case "pending":
		return "Pending"
	case "completed":
		return "Completed"
	case "cancelled":
		return "Cancelled"
	default:
		return "Pending"
	}
}

func GetIdFromParam(c *gin.Context) (int32, error) {
	orderIdStr := c.Param("id")

	if orderIdStr == "" {
		return 0, ErrInvalidRequest
	}

	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		return 0, ErrInvalidRequest
	}

	return int32(orderId), nil
}
