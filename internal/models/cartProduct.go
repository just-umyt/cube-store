package models

type CartProduct struct {
	ID        uint `gorm:"primaryKey"`
	ProductId uint
	Quantity  uint `gorm:"default:1"`
	CartId    uint
}
