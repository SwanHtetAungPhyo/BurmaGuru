package db

import (
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetUpDataBase() error {
	dsn := "root:Swanhtet12@@tcp(127.0.0.1:3306)/BurmaGuru?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Migrate() {
	err := DB.AutoMigrate(&models.Country{})
	if err != nil {
		panic("Failed to migrate the database")
	}
}
