package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/Project/internal/models"
	"github.com/shangcheng/Project/Project/internal/services"
)

type UserHandler struct {
	UserService *services.UserService
}

// 注册用户
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证注册数据
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "验证失败", "message": err.Error()})
		return
	}

	// 创建用户
	if err := h.UserService.CreateUser(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户注册成功"})
}
