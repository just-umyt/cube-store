package models

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	SessionId uint
	Products  []CartProduct
}
