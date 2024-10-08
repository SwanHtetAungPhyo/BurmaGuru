package models

import "time"

type Country struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	LivingExpense   float64   `json:"livingExpense"`
	LifeStyleRating float32   `json:"lifeStyleRating"`
	UniqueField     string    `json:"uniqueField"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
