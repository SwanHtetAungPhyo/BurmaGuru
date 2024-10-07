package models

type User struct {
	Id int64 `json:"Id"`
	UserName string `json:"UserName"`
	Email string `json:"Email"`
	Password string `json:"Password"`
	ProfilePicture string `json:"ProfilePicture"`
	IsVerified bool `json:"IsVerfied"`
	Interest string `json:"Interest"`
	Country string `json:"Country"`
	EmailVerificationToken string `json:"emailTOken"`
}

func NewUser(userName, email ,password, profilePicture,Interest, VerificationToken string, isVerified bool) *User {
	return &User{
		UserName: userName,
		Email: email, 
		Password:  password,
		ProfilePicture:  profilePicture,
		IsVerified: isVerified,
		EmailVerificationToken: VerificationToken,
	}
}

