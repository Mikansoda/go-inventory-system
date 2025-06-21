package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"day23/controllers"
)

// Jika ada CRUD request ke /orders, panggil fungsi yang sesuai dari controllers
func OrderRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/orders", controllers.GetOrders(db))
	r.POST("/orders", controllers.CreateOrder(db))
	r.PUT("/orders/:id", controllers.UpdateOrder(db))
	r.DELETE("/orders/:id", controllers.DeleteOrder(db))
}