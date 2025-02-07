package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"academic-system/internal/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dbName := os.Getenv("DB_NAME")

	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate the models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Class{},
		&models.Subject{},
		&models.Assessment{},
		&models.StudentAssessment{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database connected and migrated successfully")
}