package models

import "gorm.io/gorm"

type University struct {
	gorm.Model
	Name            string  `json:"name"`
	CountryID       uint    `json:"country_id"`
	CountryName     string  `json:"country_name"`
	Ranking         int     `json:"ranking"`
	StudentCount    int     `json:"student_count"`
	EstablishedYear int     `json:"established_year"`
	Type            string  `json:"type"`
	Website         string  `json:"website"`
	Courses         string  `json:"courses"`
	Address         string  `json:"address"`
	Description     string  `json:"description"`
	UniqueFeature   string  `json:"unique_feature"`
	AverageTuition  float64 `json:"average_tuition"`
}
