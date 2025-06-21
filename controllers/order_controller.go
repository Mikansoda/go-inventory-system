package controllers

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"day23/models"
	"day23/services"
)

// Function utk fetching order, return data dalam format JSON jika berhasil, return error jika tidak
func GetOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get order lewat service
		orders, err := services.GetAllOrders(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Return order yang ada
		c.JSON(http.StatusOK, orders)
	}
}

// Function utk membuat order baru, menerima data JSON dari request body, dan return order yang dibuat, return error jika tidak
func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		// Binding JSON dari body ke struct order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
        // Setting tanggal pemesanan dengan waktu saat ini, field OrderDate menjadi waktu sekarang
		order.OrderDate = time.Now()
		// Simpan order/create order yang dibuat ke database lewat service
		if err := services.CreateOrder(db, &order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return order yang berhasil dibuat
		c.JSON(http.StatusCreated, order)
	}
}

// Function utk update order, return data update dalam format JSON jika berhasil, return error jika tidak
func UpdateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// Cari order berdasarkan ID
		order, err := services.GetOrderByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		// Binding data JSON baru ke struct inventory yg udah ditemukan
		if err := c.ShouldBindJSON(order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Simpan order yang sudah diupdate ke database lewat service
		if err := services.UpdateOrder(db, order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
       	// Return order yang berhasil di update
		c.JSON(http.StatusOK, order)
	}
}

// Function utk delete order, return messaage jika berhasil, return error jika tidak
func DeleteOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Hapus order berdasarkan ID lewat service
		if err := services.DeleteOrder(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Return message jika berhasil
		c.JSON(http.StatusOK, gin.H{"message": "Inventory sucessfully deleted"})
	}
}

