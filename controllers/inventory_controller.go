package controllers

import (
	"day23/models"
	"day23/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Function utk fetching inventory, return data dalam format JSON jika berhasil, return error jika tidak
func GetInventories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get inventory lewat service
		inventories, err := services.GetAllInventories(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Return inventory yang ada
		c.JSON(http.StatusOK, inventories)
	}
}

// Function utk membuat inventory baru, menerima data JSON dari request body, dan return inventory yang dibuat, return error jika tidak
func CreateInventory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var inventory models.Inventory
		// Binding JSON dari body ke struct Inventory
		if err := c.ShouldBindJSON(&inventory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Simpan inventory yang dibuat ke database lewat service
		if err := services.CreateInventory(db, &inventory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return inventory yang berhasil dibuat
		c.JSON(http.StatusCreated, inventory)
	}
}

// Function utk update inventory, return data dalam format JSON jika berhasil, return error jika tidak
func UpdateInventory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Cari inventory berdasarkan ID
		inventory, err := services.GetInventoryByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
			return
		}

		// Binding data JSON baru ke struct inventory yg udah ditemukan
		if err := c.ShouldBindJSON(inventory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update inventory di database lewat service
		if err := services.UpdateInventory(db, inventory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return inventory yang berhasil di update
		c.JSON(http.StatusOK, inventory)
	}
}

// Function utk update inventory, return messaage jika berhasil, return error jika tidak
func DeleteInventory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// Hapus inventory berdasarkan ID lewat service
		if err := services.DeleteInventory(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return message jika berhasil
		c.JSON(http.StatusOK, gin.H{"message": "Inventory sucessfully deleted"})
	}
}
