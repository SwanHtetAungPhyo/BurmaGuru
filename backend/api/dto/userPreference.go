package dto

type UserPreference struct {
	ExpectedLivingCost         float64 `json:"expectedLivingCost"`
	ExpectedLifeStyleRate      float32 `json:"expectedLifeStyleRate"`
	ExpectedAverageTemperature float64 `json:"expectedAverageTemperature"`
	ReasonToMove               string  `json:"reasonToMove"`
}

type RecommendationFlag struct {
	Country     string  `json:"country"`
	Possibility float64 `json:"possibilityToMove"`
}

type RecommendationArray struct {
	Recommendation []RecommendationFlag `json:"recommendations"`
}
