package models

// 	Representasi baris data dalam tabel products, diolah lewat ORM GORM
// Product memiliki relasi:
// - One to Many dengan Inventory (satu produk bisa punya banyak inventory)
// - One to Many dengan Order (satu produk bisa ada di banyak order)
type Product struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	Category    string `json:"category"`

// Relasi
	Inventories []Inventory `gorm:"foreignKey:ProductID" json:"inventories,omitempty"`
	Orders      []Order     `gorm:"foreignKey:ProductID" json:"orders,omitempty"`
}