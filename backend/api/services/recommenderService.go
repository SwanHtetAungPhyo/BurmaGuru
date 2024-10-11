package services

import (
	"github.com/swanhtetaungphyo/burmaGuru/dto"
	"log"
	"math"

	"github.com/swanhtetaungphyo/burmaGuru/db"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
)

func Recommender(preference dto.UserPreference) ([]dto.RecommendationFlag, error) {
	var recommendedCountries []models.Country
	var recommendations []dto.RecommendationFlag

	if err := db.DB.Find(&recommendedCountries).Error; err != nil {
		log.Printf("Error in %v is %v", utils.CurrentFunction(), err.Error())
		return nil, err
	}

	for _, country := range recommendedCountries {
		possibility := calculatePossibility(preference, country)
		if possibility >= 60.0 { // Include only countries with possibility score above 30%
			recommendations = append(recommendations, dto.RecommendationFlag{
				Country:     country.Name,
				Possibility: possibility,
			})
		}
	}

	return recommendations, nil
}

func calculatePossibility(preference dto.UserPreference, country models.Country) float64 {
	// Set default weights
	lifeStyleWeight := 0.3
	livingExpenseWeight := 0.5
	averageTempWeight := 0.2
	safetyWeight := 0.15
	healthcareWeight := 0.15

	// Adjust weights based on ReasonToMove
	adjustWeights(&lifeStyleWeight, &livingExpenseWeight, &averageTempWeight, &safetyWeight, &healthcareWeight, preference.ReasonToMove)

	// Compute individual scores, handling division by zero
	lifestyleScore := safeDivision(float64(country.LifeStyleRating), float64(preference.ExpectedLifeStyleRate)) * lifeStyleWeight
	expenseScore := safeDivision(preference.ExpectedLivingCost-country.LivingExpense, preference.ExpectedLivingCost) * livingExpenseWeight
	temperatureScore := safeDivision(country.AverageTemperature, preference.ExpectedAverageTemperature) * averageTempWeight
	safetyScore := safeDivision(float64(country.SafetyIndex), float64(preference.PreferredSafetyIndex)) * safetyWeight
	healthcareScore := safeDivision(float64(country.HealthcareIndex), float64(preference.PreferredHealthcareIndex)) * healthcareWeight

	// Sum up the scores
	possibility := lifestyleScore + expenseScore + temperatureScore + safetyScore + healthcareScore

	// Normalize and return possibility as percentage
	possibility = math.Min(math.Max(possibility, 0), 1) // Ensure value is between 0 and 1
	return possibility * 100
}

func adjustWeights(lifeStyleWeight, livingExpenseWeight, averageTempWeight, safetyWeight, healthcareWeight *float64, reason string) {
	switch reason {
	case "Education":
		*livingExpenseWeight = 0.4
		*healthcareWeight = 0.25
	case "Employment":
		*livingExpenseWeight = 0.4
		*safetyWeight = 0.3
	case "Retirement":
		*lifeStyleWeight = 0.5
		*safetyWeight = 0.3
	}
}

// safeDivision performs division and handles division by zero.
func safeDivision(numerator, denominator float64) float64 {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
