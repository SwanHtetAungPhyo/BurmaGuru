package services

import (
	"github.com/swanhtetaungphyo/burmaGuru/db"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"log"
)

func CreateCountry() {

}

func GetAllCountry() ([]models.Country, error) {
	var CountryArray []models.Country
	if err := db.DB.Find(&CountryArray).Error; err != nil {
		log.Printf("Error is in %v : %v", utils.CurrentFunction(), err.Error())
		return nil, err
	}
	return CountryArray, nil
}

func GetCountryById() {

}
