package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitRoutes inisialisasi semua endpoint API yang ada
func InitRoutes(r *gin.Engine, db *gorm.DB) {
	ProductRoutes(r, db)
	InventoryRoutes(r, db)
	OrderRoutes(r, db)
	ExternalRoutes(r)
}