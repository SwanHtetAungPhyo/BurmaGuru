package services

import (
	"errors"

	"github.com/swanhtetaungphyo/burmaguru/models"
	"github.com/swanhtetaungphyo/burmaguru/utils"
)


var UserArray []models.User
var nextId = int64(1)

func RegistrationServices(user * models.User) (*models.User, error){
	
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

	user.Id = nextId
	nextId++

	UserArray = append(UserArray, *user)
	return user, nil
}