package dto

type UserDto struct {
	UserName       string `json:"UserName"`
	Email          string `json:"Email"`
	Password       string `json:"Password"`
	Interest       string `json:"Interest"`
	Country        string `json:"Country"`
	IsVerified     bool   `json:"is_verified"`
	ProfilePicture string `json:"ProfilePicture"`
}
