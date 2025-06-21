package routes

import (
	"github.com/gin-gonic/gin"
	"day23/controllers"
)

// Jika ada GET request ke /external-products, panggil fungsi FetchExternalProducts dari controllers
func ExternalRoutes(r *gin.Engine) {
	r.GET("/external-products", controllers.FetchExternalProducts)
}