package main

import (
	"github.com/gin-gonic/gin"
	"day23/config"
	"day23/routes"
	"day23/models"
)

func main() {
	db := config.ConnectDatabase()
	db.AutoMigrate(&models.Product{}, &models.Inventory{}, &models.Order{})

	r := gin.Default()
	routes.InitRoutes(r, db)

	r.Run(":8080")
}
