package models

import "time"

// 	Representasi baris data dalam tabel orders, diolah lewat ORM GORM
// Product memiliki relasi:
// - Many to One ke Product (satu order hanya berhubungan dengan satu produk)
type Order struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	OrderDate time.Time `json:"order_date"`

// Relasi
	Product Product `gorm:"foreignKey:ProductID" json:"-"`
}