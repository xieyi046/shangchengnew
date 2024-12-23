package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 用户注册路由
	SetupUserRouter(r, db)
	SetupOrderRouter(r, db)
	SetupProductRouter(r, db)
	SetupPayRouter(r, db)
	SetupPaysRouter(r, db)

	return r
}
