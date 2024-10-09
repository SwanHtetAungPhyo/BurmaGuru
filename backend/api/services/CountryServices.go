package services

import (
	"github.com/swanhtetaungphyo/burmaGuru/db"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"log"
)

func CreateCountry(country models.Country) error {
	if err := db.DB.Create(&country).Error; err != nil {
		log.Printf("Error in %v: %v", utils.CurrentFunction(), err.Error())
		return err
	}
	return nil
}

func GetAllCountry() ([]models.Country, error) {
	var countries []models.Country
	if err := db.DB.Find(&countries).Error; err != nil {
		log.Printf("Error in %v: %v", utils.CurrentFunction(), err.Error())
		return nil, err
	}
	return countries, nil
}

func GetCountryById(id uint) (models.Country, error) {
	var country models.Country
	if err := db.DB.First(&country, id).Error; err != nil {
		log.Printf("Error in %v: %v", utils.CurrentFunction(), err.Error())
		return models.Country{}, err
	}
	return country, nil
}

func UpdateCountry(id uint, updatedCountry models.Country) error {
	if err := db.DB.Model(&models.Country{}).Where("id = ?", id).Updates(updatedCountry).Error; err != nil {
		log.Printf("Error in %v: %v", utils.CurrentFunction(), err.Error())
		return err
	}
	return nil
}

func DeleteCountry(id uint) error {
	if err := db.DB.Delete(&models.Country{}, id).Error; err != nil {
		log.Printf("Error in %v: %v", utils.CurrentFunction(), err.Error())
		return err
	}
	return nil
}
