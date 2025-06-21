package services

import (
	"gorm.io/gorm"
	"day23/models"
)

// Service functions untuk operasi CRUD orders
// Semua fungsi ini ineraksi dengan database menggunakan GORM.
// Dipakai oleh controller untuk memisahkan logic database dari layer HTTP
func GetAllOrders(db *gorm.DB) ([]models.Order, error) {
	var orders []models.Order
	err := db.Find(&orders).Error
	return orders, err
}

func CreateOrder(db *gorm.DB, order *models.Order) error {
	return db.Create(order).Error
}

func UpdateOrder(db *gorm.DB, order *models.Order) error {
	return db.Save(order).Error
}

func DeleteOrder(db *gorm.DB, id string) error {
	return db.Delete(&models.Order{}, id).Error
}

func GetOrderByID(db *gorm.DB, id string) (*models.Order, error) {
	var order models.Order
	err := db.First(&order, id).Error
	return &order, err
}