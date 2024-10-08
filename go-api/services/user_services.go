package services

import (
	"database/sql"
	"errors"
	"log"

	"github.com/swanhtetaungphyo/burmaguru/databases"
	"github.com/swanhtetaungphyo/burmaguru/dto"
	"github.com/swanhtetaungphyo/burmaguru/models"
)

func GetAllUser() ([]models.User, error) {
	if databases.DataBase == nil {
		log.Println("Error: Database connection is nil")
		return nil, sql.ErrNoRows
	}

	rows, err := databases.DataBase.Query("SELECT id, username, email, password, interests, is_verified, country FROM Users")
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var usersArray []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password, &user.Interest, &user.IsVerified, &user.Country); err != nil {
			log.Printf("Error scanning user: %v", err)
			return nil, err
		}
		usersArray = append(usersArray, user)
	}

	return usersArray, nil
}

func GetUserByIdService(Id int64) (dto.UserDto, error) {
	if databases.DataBase == nil {
		log.Println("Error: Database connection is nil")
		return dto.UserDto{}, sql.ErrConnDone
	}

	var user models.User
	query := "SELECT id, username, email, password, interests, is_verified, country FROM Users WHERE id = ?"
	err := databases.DataBase.QueryRow(query, Id).Scan(&user.Id, &user.UserName, &user.Email, &user.Password, &user.Interest, &user.IsVerified, &user.Country)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User with ID %d not found", Id)
			return dto.UserDto{}, err
		}
		log.Printf("Error querying user by ID: %v", err)
		return dto.UserDto{}, err
	}

	return dto.UserDto{
		UserName:   user.UserName,
		Email:      user.Email,
		Interest:   user.Interest,
		IsVerified: user.IsVerified,
		Country:    user.Country,
	}, nil
}

func UpdateUserService(Id int64, updateUserDto dto.UserDto) (dto.UserDto, error) {
	if databases.DataBase == nil {
		log.Println("Error: Database connection is nil")
		return dto.UserDto{}, sql.ErrConnDone
	}

	query := `UPDATE Users SET username = ?, email = ?, password = ?, interests = ?, is_verified = ?, country = ? WHERE id = ?`
	_, err := databases.DataBase.Exec(query, updateUserDto.UserName, updateUserDto.Email, updateUserDto.Password, updateUserDto.Interest, updateUserDto.IsVerified, updateUserDto.Country, Id)
	if err != nil {
		log.Printf("Error updating user with ID %d: %v", Id, err)
		return dto.UserDto{}, err
	}

	return GetUserByIdService(Id)
}

func DeleteUserService(Id int64) (dto.UserDto, error) {
	if databases.DataBase == nil {
		log.Println("Error: Database connection is nil")
		return dto.UserDto{}, sql.ErrConnDone
	}

	user, err := GetUserByIdService(Id)
	if err != nil {
		return dto.UserDto{}, err
	}

	query := "DELETE FROM Users WHERE id = ?"
	result, err := databases.DataBase.Exec(query, Id)
	if err != nil {
		log.Printf("Error deleting user with ID %d: %v", Id, err)
		return dto.UserDto{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting affected rows for delete operation: %v", err)
		return dto.UserDto{}, err
	}

	if rowsAffected == 0 {
		return dto.UserDto{}, errors.New("no rows affected, user might not exist")
	}

	log.Printf("User with ID %d successfully deleted", Id)
	return user, nil
}
