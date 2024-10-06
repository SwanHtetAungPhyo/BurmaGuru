package models

type User struct {
	Id int64 `json:"Id"`
	UserName string `json:"UserName"`
	Email string `json:"Email"`
	Password string `json:"Password"`
	ProfilePicture string `json:"ProfilePicture"`
	IsVerified bool `json:"IsVerfied"`
	EmailVerificationToken string `json:"emailTOken"`
}