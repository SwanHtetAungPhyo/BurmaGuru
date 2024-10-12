package services

import (
	"fmt"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/joho/godotenv"
	"github.com/swanhtetaungphyo/burmaGuru/db"
	"github.com/swanhtetaungphyo/burmaGuru/dto"
	"github.com/swanhtetaungphyo/burmaGuru/models"
	"github.com/swanhtetaungphyo/burmaGuru/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net"
	"os"
)

const role = "USER"

// GetUser retrieves a user by their ID.
// It returns a UserDTO and an error if the user is not found or if there is a database error.
func GetUser(id int) (dto.UserDTO, error) {
	var user models.User
	var responseUser dto.UserDTO
	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		log.Fatalf("%v : %v", utils.CurrentFunction(), err)
		return responseUser, err
	}
	responseUser = utils.ToUserDTO(user)
	return responseUser, nil
}

// CreateUser creates a new user from the provided UserDTO.
// It hashes the password and fetches the user's country based on their IP address.
// Returns an error if the country cannot be determined or if there is a database error.
func CreateUser(user dto.UserDTO) error {
	var userToCreate models.User
	userToCreate = utils.ToUser(user)

	countryName := findUserLocation(user.Ip)
	userToCreate.Country = countryName

	if countryName == "" {
		log.Printf("No country name found for IP: %s", user.Ip)
		return fmt.Errorf("country name not found for the provided IP")
	}

	userToCreate.Password, _ = hashPassword(user.Password)
	userToCreate.Role = role
	if err := db.DB.Create(&userToCreate).Error; err != nil {
		log.Fatalf("%v : %v", utils.CurrentFunction(), err)
		return err
	}
	return nil
}

// DeleteUser removes a user by their ID.
// Returns an error if the deletion fails.
func DeleteUser(id int) error {
	return db.DB.Model(&models.User{}).Where("id = ?", id).Delete(&models.User{}).Error
}

// findUserLocation retrieves the country name based on the provided IP address.
// It uses the ipinfo API to fetch the information and returns the country name or an empty string if not found.
func findUserLocation(ipAddress net.IP) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%v : %v", utils.CurrentFunction(), err)
		return ""
	}
	apikey := os.Getenv("TOKEN")
	client := ipinfo.NewClient(nil, nil, apikey)
	info, err := client.GetIPInfo(net.ParseIP(string(ipAddress)))
	if err != nil {
		log.Fatalf("%v : %v", utils.CurrentFunction(), err)
		return ""
	}
	if info.CountryName == "" {
		log.Printf("No country name found for IP: %s", ipAddress)
		return ""
	}
	return info.CountryName
}

// hashPassword hashes the provided plain password using bcrypt.
// It returns the hashed password as a string and an error if the hashing fails.
func hashPassword(plainPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
