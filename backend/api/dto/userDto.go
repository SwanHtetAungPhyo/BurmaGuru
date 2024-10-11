package dto

import "net"

type UserDTO struct {
	UserName               string `json:"user_name" validator:"required,min=4,max=30"`
	Email                  string `json:"email" validator:"required,email"`
	Ip                     net.IP `json:"ip" validator:"required"`
	Password               string `json:"password" validator:"required,min=8"`
	ProfilePicture         string `json:"profile_picture"`
	IsVerified             bool   `json:"is_verified"`
	Interest               string `json:"interest" validator:"required"`
	Country                string `json:"country,omitempty"`
	EmailVerificationToken string `json:"email_verification_token,omitempty"`
}
