package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/internal/models"
	"github.com/shangcheng/Project/internal/services"
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

// 用户登录
func (h *UserHandler) LoginUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用 service 层进行登录验证
	_, _, err := h.UserService.Login(req.UserName, req.PassWord)
	if err != nil {
		// 错误处理
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 登录成功，返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
	})
}

// 更新用户信息
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证更新的数据
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "验证失败", "message": err.Error()})
		return
	}

	// 调用 service 层进行更新
	if err := h.UserService.UpdateUser(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功"})
}
