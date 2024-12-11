package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/Project/handlers"
	"github.com/shangcheng/Project/Project/internal/services"
	"github.com/shangcheng/Project/Project/internal/dao"
	"gorm.io/gorm"
)

func SetupProductRouter(r *gin.Engine, db *gorm.DB) {
	// 创建 OrderDao 和 OrderService
	productDao := &dao.ProductDao{DB: db}
	productService := &services.ProductService{ProductDao: productDao}
	productHandler := &handlers.ProductHandler{ProductService: productService}

	productRoutes := r.Group("/products")
	{
		// 添加产品
		productRoutes.POST("/", productHandler.AddProduct)

		// 获取所有产品，支持分页
		productRoutes.GET("/", productHandler.GetAllProducts)

		// 获取单个产品，传入ID作为参数
		productRoutes.GET("/:id", productHandler.GetProductById)

		// 更新产品信息
		productRoutes.PUT("/:id", productHandler.UpdateProduct)

		// 删除产品
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

}
