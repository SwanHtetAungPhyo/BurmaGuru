package models

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Title     string  `gorm:"not null" json:"title,omitempty"`
	Content   string  `gorm:"not null" json:"Content,omitempty"`
	CountryId uint    `gorm:"not null" json:"country_id"`
	Link      *string `gorm:"default:null" json:"link,omitempty"`
	Photo     *string `gorm:"default:null" json:"Photo,omitempty"`
}
