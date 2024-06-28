package database

import (
	"log"
	"os"

	"github.com/just-umyt/cube-store/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := os.Getenv("DB")

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	DB.AutoMigrate(&models.Category{}, &models.Product{}, &models.User{}, &models.Cart{}, &models.Session{}, &models.CartProduct{})

}
