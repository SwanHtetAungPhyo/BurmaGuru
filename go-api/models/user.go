package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName               string `json:"user_name"`
	Email                  string `json:"email"`
	Password               string `json:"password"`
	ProfilePicture         string `json:"profile_picture"`
	IsVerified             bool   `json:"is_verified"`
	Interest               string `json:"interest"`
	Country                string `json:"country"`
	EmailVerificationToken string `json:"email_verification_token"`
}

func NewUser(userName, email, password, profilePicture, Interest, VerificationToken string, isVerified bool) *User {
	return &User{
		UserName:               userName,
		Email:                  email,
		Password:               password,
		ProfilePicture:         profilePicture,
		IsVerified:             isVerified,
		EmailVerificationToken: VerificationToken,
	}
}
