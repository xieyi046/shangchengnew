package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/internal/models"
	"github.com/shangcheng/Project/internal/services"
)

type PaysHandler struct {
	PaysService *services.PaysService
}

func NewPaysHandler(paysService *services.PaysService) *PaysHandler {
	return &PaysHandler{PaysService: paysService}
}

func (h *PaysHandler) MultiPayment(c *gin.Context) {
	var req models.MultiPay
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体"})
		return
	}

	err := h.PaysService.MultiPayOrder(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "支付成功"})
}
