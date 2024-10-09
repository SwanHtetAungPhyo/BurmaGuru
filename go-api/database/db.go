package database

import (
	"github.com/swanhtetaungphyo/burmaguru/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global database instance

// SetUpDataBase initializes the database connection and sets the global DB instance
func SetUpDataBase() error {
	dsn := "root:Swanhtet12@@tcp(127.0.0.1:3306)/BurmaGuru?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

// Migrate handles database migrations
func Migrate() {
	err := DB.AutoMigrate(&models.User{}) // Adjust as needed for your models
	if err != nil {
		panic("Failed to migrate the database")
	}
}
