package db

import (
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetUpDataBase() error {
	dsn := "root:Swanhtet12@@tcp(127.0.0.1:3306)/MVPOne?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Migrate() {
	migrationArray := []interface{}{
		&models.Country{}, &models.University{},
		&models.Resource{},
	}

	for _, table := range migrationArray {
		err := DB.AutoMigrate(table)
		if err != nil {
			panic("Failed to migrate the database")
		}
	}

}
