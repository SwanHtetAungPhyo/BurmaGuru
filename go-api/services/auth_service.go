package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/swanhtetaungphyo/burmaguru/database"
	"github.com/swanhtetaungphyo/burmaguru/dto"
	"github.com/swanhtetaungphyo/burmaguru/models"
	"github.com/swanhtetaungphyo/burmaguru/utils"
	"golang.org/x/crypto/bcrypt"
)

var UserArray []models.User

func RegistrationServices(user *dto.UserDto) (*models.User, error) {
	if strings.TrimSpace(user.UserName) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		return nil, errors.New("userName, Email, and Password must be filled")
	}

	hashedPassword, err := utils.HashPassowrd(user.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	insertSql, err := utils.LoadSqlFile("insert_user.sql")
	if err != nil {
		return nil, errors.New("failed to load SQL file")
	}

	token := utils.EmailTokenGenerator(10)
	newUser := models.NewUser(user.UserName, user.Email, user.Password, user.ProfilePicture, user.Interest, token, false)

	if database.DataBase == nil {
		return nil, errors.New("database connection is not initialized")
	}

	_, err = database.DataBase.Exec(insertSql, newUser.UserName, newUser.Email, newUser.Password, newUser.ProfilePicture, newUser.EmailVerificationToken, newUser.IsVerified)
	if err != nil {
		return nil, errors.New("failed to register user")
	}

	err = SendVerificationEmail(newUser.Email, token)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func LoginService(email string, password string) error {
	existingUser := FindUserByEmail(email)
	if existingUser.Email == "" {
		return fmt.Errorf("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
}

func FindUserByEmail(email string) models.User {
	for _, existingUser := range UserArray {
		if existingUser.Email == email {
			return existingUser
		}
	}
	return models.User{}
}

func SendVerificationEmail(email, token string) error {
	subject := "Email Verification"
	body := "Click the following link to verify your email: " +
		"http://localhost:8080/verify-email?token=" + token
	return utils.SendEmail([]string{email}, subject, body)
}
