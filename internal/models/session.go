package models

type Session struct {
	ID     uint `gorm:"primaryKey"`
	UserId *uint
	Cart   Cart
}
