package services

import (
	"github.com/swanhtetaungphyo/burmaGuru/db"
	"github.com/swanhtetaungphyo/burmaGuru/dto"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"log"
)

const (
	lifeStyleWeight     = 0.3
	livingExpenseWeight = 0.5
	averageTempWeight   = 0.2
	minTemp             = -20.0
	maxTemp             = 20
)

func Recommender(preference dto.UserPreference) ([]dto.RecommendationFlag, interface{}) {
	var recommendedCountries []models.Country
	var recommendations []dto.RecommendationFlag

	if err := db.DB.Find(&recommendedCountries).Error; err != nil {
		log.Printf("Error in %v is %v", utils.CurrentFunction(), err.Error())
		return nil, err
	}

	for _, country := range recommendedCountries {
		possibilty := calculatePossibility(preference, country)
		if possibilty >= 30.000 {
			recommendations = append(recommendations, dto.RecommendationFlag{
				Country:     country.Name,
				Possibility: possibilty,
			})
		}

	}
	return recommendations, nil
}

func calculatePossibility(preference dto.UserPreference, country models.Country) float64 {
	lifestyleScore := (float64(country.LifeStyleRating) / float64(preference.ExpectedLifeStyleRate)) * lifeStyleWeight
	expenseScore := ((preference.ExpectedLivingCost - country.LivingExpense) / preference.ExpectedLivingCost) * livingExpenseWeight
	temperatureScore := (country.AverageTemperature / preference.ExpectedAverageTemperature) * averageTempWeight

	possibility := lifestyleScore + expenseScore + temperatureScore
	if possibility < 0 {
		return 0
	}
	if possibility > 1 {
		return 1
	}
	return possibility * 100
}
