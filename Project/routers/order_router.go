package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/Project/handlers"
	"github.com/shangcheng/Project/Project/internal/services"
	"github.com/shangcheng/Project/Project/internal/dao"
	"gorm.io/gorm"
)

func SetupOrderRouter(r *gin.Engine, db *gorm.DB) {
	// 创建 OrderDao 和 OrderService
	orderDao := &dao.OrderDao{DB: db}
	orderService := &services.OrderService{OrderDao: orderDao}
	orderHandler := &handlers.OrderHandler{OrderService: orderService}

	orderGroup := r.Group("/orders")
	{
		// 创建订单
		orderGroup.POST("/", orderHandler.CreateOrder)

		// 查询订单
		orderGroup.GET("/:order_id", orderHandler.GetOrder)

		// 删除订单
		orderGroup.DELETE("/:order_id", orderHandler.DeleteOrder)

		// 获取订单详情
		orderGroup.GET("/:order_id/details", orderHandler.GetOrderDetails)
	}
}
