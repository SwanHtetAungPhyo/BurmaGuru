package dto

type UserDto struct {
	UserName string `json:"UserName"`
	Email string `json:"Email"`
	Password string `json:"Password"`
	Interest string `json:"Interest"`
	ProfilePicture string `json:"ProfilePicture"`
}