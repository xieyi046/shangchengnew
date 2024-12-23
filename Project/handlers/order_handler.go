package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/Project/internal/models"
	"github.com/shangcheng/Project/Project/internal/services"

	"strconv"
)

// OrderHandler 处理订单的请求
type OrderHandler struct {
	OrderService *services.OrderService
}

// 创建订单
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	// 解析请求的 JSON 数据
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	userId := c.GetInt("user_id") // 假设用户ID存储在上下文中，或者通过请求传递
	order.UserId = userId         // 设置用户ID

	// 调用服务层的创建订单方法
	if err := h.OrderService.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "订单创建成功",
		"order":   order,
	})
}

// 订单列表
func (h *OrderHandler) GetOrderList(c *gin.Context) {
	userId := c.GetInt("user_id") // 假设用户ID存储在上下文中，或者通过请求传递

	// 获取分页参数 (可选)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	orders, err := h.OrderService.GetOrdersByUserId(userId, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
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
