package models

import "gorm.io/gorm"

type University struct {
	gorm.Model
	Name            string  `json:"name"`             // University name
	CountryID       uint    `json:"country_id"`       // Reference to Country ID
	CountryName     string  `json:"country_name"`     // Country name (for easier querying)
	Ranking         int     `json:"ranking"`          // Global ranking of the university
	StudentCount    int     `json:"student_count"`    // Number of enrolled students
	EstablishedYear int     `json:"established_year"` // Year the university was established
	Type            string  `json:"type"`             // Type of university (e.g., Public, Private)
	Website         string  `json:"website"`          // Official university website link
	Courses         string  `json:"courses"`          // Popular courses offered by the university
	Address         string  `json:"address"`          // Physical address of the university
	Description     string  `json:"description"`      // Brief description of the university
	UniqueFeature   string  `json:"unique_feature"`   // Outstanding feature (e.g., Research, Technology, Business)
	AverageTuition  float64 `json:"average_tuition"`  // Average annual tuition fees in USD
}
