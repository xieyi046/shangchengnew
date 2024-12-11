package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/Project/internal/services"

	"strconv"
)

// OrderHandler 处理订单的请求
type OrderHandler struct {
	OrderService *services.OrderService
}

// 创建订单
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var request struct {
		UserId     int     `json:"user_id"`
		OrderPrice float64 `json:"order_price"`
		PayType    string  `json:"pay_type"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.OrderService.CreateOrder(request.UserId, request.OrderPrice, request.PayType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "订单创建成功",
		"order":   order,
	})
}

// 查询订单
func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderIdStr := c.Param("order_id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的订单ID"})
		return
	}

	order, err := h.OrderService.GetOrderById(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

// 删除订单
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	orderIdStr := c.Param("order_id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的订单ID"})
		return
	}

	userId := c.GetInt("user_id") // 假设用户ID存储在上下文中，或者可以通过请求传递

	err = h.OrderService.DeleteOrder(orderId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "订单删除成功",
	})
}

// 获取订单详情
func (h *OrderHandler) GetOrderDetails(c *gin.Context) {
	orderIdStr := c.Param("order_id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的订单ID"})
		return
	}

	order, err := h.OrderService.GetOrderById(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

