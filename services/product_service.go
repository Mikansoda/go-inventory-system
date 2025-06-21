package services

import (
	"gorm.io/gorm"
	"day23/models"
)

// Service functions untuk operasi CRUD products
// Semua fungsi ini ineraksi dengan database menggunakan GORM.
// Dipakai oleh controller untuk memisahkan logic database dari layer HTTP
func GetAllProducts(db *gorm.DB) ([]models.Product, error) {
	var products []models.Product
	err := db.Find(&products).Error
	return products, err
}

func CreateProduct(db *gorm.DB, product *models.Product) error {
	return db.Create(product).Error
}

func UpdateProduct(db *gorm.DB, product *models.Product) error {
	return db.Save(product).Error
}

func DeleteProduct(db *gorm.DB, id string) error {
	return db.Delete(&models.Product{}, id).Error
}

func GetProductByID(db *gorm.DB, id string) (*models.Product, error) {
	var product models.Product
	err := db.First(&product, id).Error
	return &product, err
}