package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAdminResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (h *UserHandler) GetAdmins(c *gin.Context) {
	admins, err := h.service.GetAdmins(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": ErrServer})
		return
	}

	var response []GetAdminResponse
	for _, admin := range admins {
		response = append(response, GetAdminResponse{
			ID:    int64(admin.ID),
			Email: admin.Email,
			Role:  admin.Role,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admins retrieved successfully",
		"admins":  response,
	})
}
