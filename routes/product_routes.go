package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"day23/controllers"
)

// Jika ada CRUD request ke /products, panggil fungsi yang sesuai dari controllers
func ProductRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/products", controllers.GetProducts(db))
	r.POST("/products", controllers.CreateProduct(db))
	r.PUT("/products/:id", controllers.UpdateProduct(db))
	r.DELETE("/products/:id", controllers.DeleteProduct(db))
}