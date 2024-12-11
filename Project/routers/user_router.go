package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/Project/handlers"
	"github.com/shangcheng/Project/Project/internal/services"
	"github.com/shangcheng/Project/Project/internal/dao"
	"gorm.io/gorm"
)

func SetupUserRouter(r *gin.Engine, db *gorm.DB) {
	// 创建 UserDao 和 UserService
	userDao := &dao.UserDao{DB: db}
	userService := &services.UserService{UserDao: userDao}
	userHandler := &handlers.UserHandler{UserService: userService}

	userRoutes := r.Group("/user")
	{
		// 注册用户
		userRoutes.POST("/register", userHandler.RegisterUser)
	}
}