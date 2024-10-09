package services

import (
	"github.com/swanhtetaungphyo/burmaguru/database"
	"github.com/swanhtetaungphyo/burmaguru/dto"
	"github.com/swanhtetaungphyo/burmaguru/models"
	"gorm.io/gorm"
	"log"
)

func GetAllUser(db *gorm.DB) ([]models.User, error) {
	if db == nil {
		log.Println("Error: Database connection is nil")
		return nil, gorm.ErrInvalidDB
	}

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}

	return users, nil
}

func GetUserByIdService(Id int64) (dto.UserDto, error) {
	var user models.User
	result := database.DB.First(&user, Id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return dto.UserDto{}, result.Error
		}
	}

	resultUser := &dto.UserDto{
		UserName:       user.UserName,
		Email:          user.Email,
		Password:       user.Password,
		ProfilePicture: user.ProfilePicture,
		IsVerified:     user.IsVerified,
		Interest:       user.Interest,
		Country:        user.Country,
	}
	return *resultUser, nil
}

func UpdateUserService(email string, updatedUser *dto.UserDto) (models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	user.UserName = updatedUser.UserName
	user.Password = updatedUser.Password
	user.ProfilePicture = updatedUser.ProfilePicture
	user.IsVerified = updatedUser.IsVerified
	user.Interest = updatedUser.Interest
	user.Country = updatedUser.Country
	return user, nil
}

func DeleteUserService(Id int64) error {
	return database.DB.Delete(&models.User{}, Id).Error
}
