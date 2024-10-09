package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name            string  `json:"name"`
	LivingExpense   float64 `json:"livingExpense"`
	LifeStyleRating float32 `json:"lifeStyleRating"`
	UniqueField     string  `json:"uniqueField"`
}
