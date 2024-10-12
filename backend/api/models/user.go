package models

import (
	_ "github.com/swanhtetaungphyo/burmaGuru/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName               string `gorm:"column:user_name" json:"user_name"`
	Email                  string `gorm:"unique,not null,column:email" json:"email"`
	Password               string `gorm:"column:password" json:"password"`
	Role                   string `gorm:"column:role" json:"role"`
	ProfilePicture         string `gorm:"column:profile_picture" json:"profile_picture"`
	IsVerified             bool   `gorm:"column:is_verified" json:"is_verified"`
	Interest               string `gorm:"column:interest" json:"interest"`
	Country                string `gorm:"column:country" json:"country"`
	EmailVerificationToken string `gorm:"column:email_verification_token" json:"email_verification_token"`
}
