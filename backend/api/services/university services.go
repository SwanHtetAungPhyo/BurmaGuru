package services

import (
	"github.com/swanhtetaungphyo/burmaGuru/db"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"log"
)

func CreateUniversity(university models.University) error {
	if err := db.DB.Create(&university).Error; err != nil {
		log.Printf("Error in %v: %v", utils.CurrentFunction(), err.Error())
		return err
	}
	return nil
}

func UniUpdateService(id uint, updatedData models.University) error {
	if err := db.DB.Model(&models.University{}).Where("id = ?", id).Updates(updatedData).Error; err != nil {
		log.Printf("Error in %v: %v", utils.CurrentFunction(), err.Error())
		return err
	}
	return nil
}

func GetUniversities(countryID uint) ([]models.University, error) {
	var universities []models.University
	if err := db.DB.Where("country_id = ?", countryID).Find(&universities).Error; err != nil {
		return nil, err
	}
	return universities, nil
}
