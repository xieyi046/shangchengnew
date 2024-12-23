package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/Project/handlers"
	"github.com/shangcheng/Project/Project/internal/dao"
	"github.com/shangcheng/Project/Project/internal/services"
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
		productRoutes.POST("/add", productHandler.AddProduct)

		// 获取所有产品，支持分页
		productRoutes.GET("/getall", productHandler.GetAllProducts)

		// 获取单个产品，传入ID作为参数
		productRoutes.GET("/get", productHandler.GetProductById)

		// 更新产品信息
		productRoutes.PUT("/update", productHandler.UpdateProduct)

		// 删除产品
		productRoutes.DELETE("/delete", productHandler.DeleteProduct)
	}

}
