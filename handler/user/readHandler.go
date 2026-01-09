package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid payload"})
		return
	}

	user, err := h.userReadUC.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := h.pasetoSvc.GenerateToken(
		strconv.FormatUint(user.ID, 10),
		user.Role,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(200, gin.H{
		"access_token": token,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
