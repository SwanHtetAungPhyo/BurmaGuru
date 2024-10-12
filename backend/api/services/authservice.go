package services

import (
	"errors"
	"github.com/swanhtetaungphyo/burmaGuru/db"
	"github.com/swanhtetaungphyo/burmaGuru/dto" // Import dto for LoginRequest and UserDTO
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"os"
)

func Registration(user *dto.UserDTO) error {
	return nil
}

func GenereteTokens(user models.User) (string, string, error) {
	accessToken, err := utils.GenereteTokens("Access", user.UserName, user.Role, os.Getenv("ACCESS_TOKEN_EXPIRATION"))
	if err != nil {
		return "", "", err
	}
	refreshToken, err := utils.GenereteTokens("Refresh", user.UserName, user.Role, os.Getenv("REFRESH_TOKEN_EXPIRATION"))
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func Login(loginData *dto.LoginRequest) (string, string, error) {
	var user models.User

	if err := db.DB.Where("email = ?", loginData.Username).First(&user).Error; err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if !utils.CompareHashAndPassword(user.Password, loginData.Password) {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, refreshToken, err := GenereteTokens(user)
	if err != nil {
		return "", "", errors.New("failed to generate token")
	}

	return accessToken, refreshToken, nil
}
