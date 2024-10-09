package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name               string  `json:"name" gorm:"not null;unique"`
	Capital            string  `json:"capital"`
	Continent          string  `json:"continent"`
	Currency           string  `json:"currency"`
	LivingExpense      float64 `json:"livingExpense"`
	LifeStyleRating    float32 `json:"lifeStyleRating"`
	Population         int     `json:"population"`
	Language           string  `json:"language"`
	AverageTemperature float64 `json:"averageTemperature"`
	UniqueField        string  `json:"uniqueField"`
	Description        string  `json:"description" gorm:"type:text"`
}
