package models

import "gorm.io/gorm"

type University struct {
	gorm.Model
	Name                           string  `json:"name"`
	CountryID                      uint    `json:"country_id"`
	CountryName                    string  `json:"country_name"`
	Ranking                        int     `json:"ranking"`
	StudentCount                   int     `json:"student_count"`
	EstablishedYear                int     `json:"established_year"`
	Type                           string  `json:"type"`
	Website                        string  `json:"website"`
	Courses                        string  `json:"courses"`
	Address                        string  `json:"address"`
	Description                    string  `json:"description"`
	UniqueFeature                  string  `json:"unique_feature"`
	AverageTuition                 float64 `json:"average_tuition"`
	InternationalStudentPercentage float32 `json:"international_student_percentage"`
	EmploymentRate                 float32 `json:"employment_rate"`
	ResearchOutput                 int     `json:"research_output"`
	CampusSize                     float64 `json:"campus_size"` // in square kilometers
	StudentSatisfactionScore       float32 `json:"student_satisfaction_score"`
}
