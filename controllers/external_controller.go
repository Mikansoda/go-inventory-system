package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"day23/services"
)

// Function utk fetching external products
func FetchExternalProducts(c *gin.Context) {
	// Panggil service utk mendapatkan data produk eksternal
	data, err := services.GetExternalProducts()
	// Cek apakah ada error saat mengambil data
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Kembalikan data dalam format JSON jika berhasil
	c.Data(http.StatusOK, "application/json", data)
}