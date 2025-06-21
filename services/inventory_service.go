package services

import (
	"gorm.io/gorm"
	"day23/models"
)

// Service functions untuk operasi CRUD inventories
// Semua fungsi ini ineraksi dengan database menggunakan GORM.
// Dipakai oleh controller untuk memisahkan logic database dari layer HTTP
func GetAllInventories(db *gorm.DB) ([]models.Inventory, error) {
	var inventories []models.Inventory
	err := db.Find(&inventories).Error
	return inventories, err
}

func CreateInventory(db *gorm.DB, inventory *models.Inventory) error {
	return db.Create(inventory).Error
}

func UpdateInventory(db *gorm.DB, inventory *models.Inventory) error {
	return db.Save(inventory).Error
}

func DeleteInventory(db *gorm.DB, id string) error {
	return db.Delete(&models.Inventory{}, id).Error
}

func GetInventoryByID(db *gorm.DB, id string) (*models.Inventory, error) {
	var inventory models.Inventory
	err := db.First(&inventory, id).Error
	return &inventory, err
}