package models

type Country struct {
	ID                  uint    `gorm:"primaryKey" json:"id"`
	Name                string  `gorm:"not null" json:"name"`
	Capital             string  `gorm:"not null" json:"capital"`
	Continent           string  `gorm:"not null" json:"continent"`
	Currency            string  `gorm:"not null" json:"currency"`
	LivingExpense       float64 `gorm:"not null" json:"living_expense"`
	LifeStyleRating     float64 `gorm:"not null" json:"lifestyle_rating"`
	Population          int     `gorm:"not null" json:"population"`
	Language            string  `gorm:"not null" json:"language"`
	AverageTemperature  float64 `gorm:"not null" json:"average_temperature"`
	UniqueField         string  `gorm:"not null" json:"unique_field"`
	Description         string  `gorm:"not null" json:"description"`
	SafetyIndex         float64 `gorm:"default:0" json:"safety_index"`
	HealthcareIndex     float64 `gorm:"default:0" json:"healthcare_index"`
	CostOfLivingIndex   float64 `gorm:"default:0" json:"cost_of_living_index"`
	PollutionIndex      float64 `gorm:"default:0" json:"pollution_index"`
	ClimateType         string  `gorm:"default:''" json:"climate_type"`
	EconomicFreedomRank int     `gorm:"default:0" json:"economic_freedom_rank"`
}
