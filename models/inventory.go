package models

// 	Representasi baris data dalam tabel inventories, diolah lewat ORM GORM
// Product memiliki relasi:
// - Many to One ke Product (setiap inventory berhubungan dengan satu produk)
type Inventory struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Location  string `json:"location"`

// Relasi
Product Product `gorm:"foreignKey:ProductID" json:"-"`
}