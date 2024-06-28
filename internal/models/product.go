package models

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	About       string
	Price       float32
	Bought      int
	CategoryId  uint
	CartProduct []CartProduct
}
