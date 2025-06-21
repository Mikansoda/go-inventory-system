package controllers

import (
	"day23/models"
	"day23/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Function utk fetching product, return data dalam format JSON jika berhasil, return error jika tidak
func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get product lewat service
		products, err := services.GetAllProducts(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Return product yang ada
		c.JSON(http.StatusOK, products)
	}
}

// Function utk membuat product baru, menerima data JSON dari request body, dan return product yang dibuat, return error jika tidak
func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		// Binding JSON dari body ke struct product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Simpan product yang dibuat ke database lewat service
		if err := services.CreateProduct(db, &product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Return product yang berhasil dibuat
		c.JSON(http.StatusCreated, product)
	}
}

// Function utk update product, return data update dalam format JSON jika berhasil, return error jika tidak
func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// Cari order berdasarkan ID
		product, err := services.GetProductByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		// Binding data JSON baru ke struct product yg udah ditemukan
		if err := c.ShouldBindJSON(product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Simpan product yang sudah diupdate ke database lewat service
		if err := services.UpdateProduct(db, product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Return product yang berhasil di update
		c.JSON(http.StatusOK, product)
	}
}

// Function utk delete product, return messaage jika berhasil, return error jika tidak
func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// Hapus product berdasarkan ID lewat service
		if err := services.DeleteProduct(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Return message jika berhasil
		c.JSON(http.StatusOK, gin.H{"message": "Inventory sucessfully deleted"})
	}
}
