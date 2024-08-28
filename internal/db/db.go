package db

import (
	"invoice-management/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=test dbname=ontrack port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	// DB.AutoMigrate(&invoice.Invoice{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the Invoice model
	if err := DB.AutoMigrate(&models.Invoice{}); err != nil {
		log.Fatal("Failed to auto migrate Invoice model:", err)
	}
}
