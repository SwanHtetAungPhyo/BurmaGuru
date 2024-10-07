package services

import (
	"database/sql"
	"log"

	"github.com/swanhtetaungphyo/burmaguru/databases"
	"github.com/swanhtetaungphyo/burmaguru/models"
)

func GetAllUser() ([]models.User, error) {
    if databases.DataBase == nil {
        log.Println("Error: Database connection is nil")
        return nil,sql.ErrNoRows
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
        if err := rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password,&user.Interest,&user.IsVerified,&user.Country); err != nil {
            log.Printf("Error scanning user: %v", err)
            return nil, err
        }
        usersArray = append(usersArray, user)
    }
    return usersArray, nil
}