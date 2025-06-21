package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"day23/controllers"
)

// Jika ada CRUD request ke /inventories, panggil fungsi yang sesuai dari controllers
func InventoryRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/inventories", controllers.GetInventories(db))
	r.POST("/inventories", controllers.CreateInventory(db))
	r.PUT("/inventories/:id", controllers.UpdateInventory(db))
	r.DELETE("/inventories/:id", controllers.DeleteInventory(db))
}