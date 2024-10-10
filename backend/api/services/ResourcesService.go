package services

import (
	"github.com/swanhtetaungphyo/burmaGuru/db"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"log"
)

func CreateResource(resource models.Resource) error {
	if err := db.DB.Create(&resource).Error; err != nil {
		log.Printf("%v have error \n : %v", utils.CurrentFunction(), err.Error())
		return err
	}
	return nil
}

func GetAllResources(id uint) (interface{}, interface{}) {
	var resources []models.Resource
	if err := db.DB.Where("country_id = ?", id).First(&resources).Error; err != nil {
		log.Printf("%v have error \n  : %v", utils.CurrentFunction(), err.Error())
		return resources, err
	}
	return resources, nil
}

func UpdateResourcesSer(countryId uint, resource models.Resource) interface{} {
	if err := db.DB.Model(&models.Resource{}).Where("country_id = ?", countryId).Updates(resource).Error; err != nil {
		log.Printf("%v \n have error \n  | %v", utils.CurrentFunction(), err.Error())
		return err
	}
	return nil
}

func DeleteResourceSer(countryId uint) error {

	return db.DB.Model(&models.Resource{}).Where("country_id = ?").Delete(&models.Resource{}).Error
}
