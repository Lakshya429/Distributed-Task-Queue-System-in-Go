package database

import (
	"log"

	"github.com/Lakshya429/distributed-task-queue/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB = db
	err = DB.AutoMigrate(&models.User{}, &models.Video{})

	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}
	log.Println("Database connected successfully!")
}

func GetDB() *gorm.DB {
	return DB
}
