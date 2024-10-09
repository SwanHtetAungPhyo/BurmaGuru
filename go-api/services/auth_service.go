package services

import (
	"errors"
	"github.com/swanhtetaungphyo/burmaguru/database"
	"github.com/swanhtetaungphyo/burmaguru/models"
	"strings"

	"github.com/swanhtetaungphyo/burmaguru/dto"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)

func RegistrationServices(user *dto.UserDto) error {
	var userToSave models.User

	// Check for empty fields
	if strings.TrimSpace(user.UserName) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		return errors.New("userName, Email, and Password must be filled")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassowrd(user.Password) // Fixed typo in function name
	if err != nil {
		return errors.New("failed to hash password")
	}
	userToSave.Password = hashedPassword

	userToSave.UserName = user.UserName
	userToSave.Email = user.Email
	userToSave.Interest = user.Interest
	userToSave.Country = user.Country
	userToSave.IsVerified = false
	userToSave.ProfilePicture = user.ProfilePicture

	if err := database.DB.Create(&userToSave).Error; err != nil {
		return errors.New("creation failed")
	}

	return nil
}

//
//func LoginService(email string, password string) error {
//	existingUser := FindUserByEmail(email)
//	if existingUser.Email == "" {
//		return fmt.Errorf("user not found")
//	}
//
//	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password))
//	if err != nil {
//		return fmt.Errorf("invalid password")
//	}
//
//	return nil
//}
//
//func FindUserByEmail(email string) models.User {
//	for _, existingUser := range UserArray {
//		if existingUser.Email == email {
//			return existingUser
//		}
//	}
//	return models.User{}
//}
//
//func SendVerificationEmail(email, token string) error {
//	subject := "Email Verification"
//	body := "Click the following link to verify your email: " +
//		"http://localhost:8080/verify-email?token=" + token
//	return utils.SendEmail([]string{email}, subject, body)
//}
