package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/Project/internal/models"
	"github.com/shangcheng/Project/Project/internal/services"
)

type PayHandler struct {
	PayService *services.PayService
}

func NewPayHandler(payService *services.PayService) *PayHandler {
	return &PayHandler{PayService: payService}
}

func (h *PayHandler) Payment(c *gin.Context) {
	var req models.Pay
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体"})
		return
	}

	err := h.PayService.PayOrder(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "支付成功"})
}
