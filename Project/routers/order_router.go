package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/handlers"
	"github.com/shangcheng/Project/internal/dao"
	"github.com/shangcheng/Project/internal/services"
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
		orderGroup.POST("/creat", orderHandler.CreateOrder)

		// 查询订单
		orderGroup.GET("/get", orderHandler.GetOrder)

		// 获取订单详情
		orderGroup.GET("/getdetails", orderHandler.GetOrderDetails)

		// 删除订单
		orderGroup.DELETE("/delete", orderHandler.DeleteOrder)

		// 获取订单列表
		orderGroup.GET("/list", orderHandler.GetOrderList)
	}
}
