package services

import (
	"errors"
	"fmt"

	"github.com/swanhtetaungphyo/burmaguru/dto"
	"github.com/swanhtetaungphyo/burmaguru/models"
	"github.com/swanhtetaungphyo/burmaguru/utils"
	"golang.org/x/crypto/bcrypt"
)


var UserArray []models.User
var nextId = int64(1)

func RegistrationServices(user * dto.UserDto) (*models.User, error){
	
	if user.UserName == " " || user.Email == " " || user.Password == " " {
		return nil, errors.New("userName , Email and Password must be filled")
	}
	
	hashedPassword, err := utils.HashPassowrd(user.Password);
	if  err != nil {
		return nil, errors.New("failed hash")
	}
	user.Password = hashedPassword


	for _, existingUser := range UserArray{
		if existingUser.Email == user.Email {
			return nil, errors.New("this email is already registered ")
		}
	}
	token := utils.EmailTokenGenerator(10)
	newUser := models.NewUser(nextId,user.UserName, user.Email,user.Password, user.ProfilePicture,token, false )
	nextId++
	UserArray = append(UserArray, *newUser)
	
	err = SendVerificationEmail(newUser.Email, token )
	if err != nil{
		return &models.User{}, err
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

func FindUserByEmail(email string ) models.User{
	for _, existingUser := range UserArray{
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
	return utils.SendEmail([]string{email},subject, body)
}