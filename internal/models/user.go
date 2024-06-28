package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Phone     string
	Email     string `gorm:"unique"`
	Password  string
	Adress    string
	CreatedAt time.Time
	SessionId Session
}
