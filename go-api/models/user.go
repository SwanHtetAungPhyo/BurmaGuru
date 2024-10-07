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

func NewUser(Id int64,userName, email ,password, profilePicture, VerificationToken string, isVerified bool) *User {
	return &User{
		Id: Id, 
		UserName: userName,
		Email: email, 
		Password:  password,
		ProfilePicture:  profilePicture,
		IsVerified: isVerified,
		EmailVerificationToken: VerificationToken,
	}
}

