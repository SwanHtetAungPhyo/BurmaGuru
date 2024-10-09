package services

//
//import (
//	"database/sql"
//	"github.com/swanhtetaungphyo/burmaguru/database"
//	"github.com/swanhtetaungphyo/burmaguru/models"
//	"log"
//)
//
//// CreateCategoryService creates a new category
//func CreateCategoryService(category models.Category) (models.Category, error) {
//	if database.DataBase == nil {
//		log.Println("Error: Database connection is nil")
//		return models.Category{}, sql.ErrNoRows
//	}
//
//	query := "INSERT INTO Categories (name) VALUES (?)"
//	result, err := database.DataBase.Exec(query, category.Name)
//	if err != nil {
//		log.Printf("Error inserting category: %v", err)
//		return models.Category{}, err
//	}
//
//	id, err := result.LastInsertId()
//	if err != nil {
//		log.Printf("Error getting last insert id: %v", err)
//		return models.Category{}, err
//	}
//	category.ID = int(id)
//
//	return category, nil
//}
//
//// GetCategoryByIdService retrieves a category by its ID
//func GetCategoryByIdService(id int) (models.Category, error) {
//	if database.DataBase == nil {
//		log.Println("Error: Database connection is nil")
//		return models.Category{}, sql.ErrNoRows
//	}
//
//	var category models.Category
//	query := "SELECT id, name FROM Categories WHERE id = ?"
//	err := database.DataBase.QueryRow(query, id).Scan(&category.ID, &category.Name)
//	if err != nil {
//		log.Printf("Error retrieving category: %v", err)
//		return models.Category{}, err
//	}
//
//	return category, nil
//}
//
//// UpdateCategoryService updates an existing category
//func UpdateCategoryService(category models.Category) error {
//	if database.DataBase == nil {
//		log.Println("Error: Database connection is nil")
//		return sql.ErrNoRows
//	}
//
//	query := "UPDATE Categories SET name = ? WHERE id = ?"
//	_, err := database.DataBase.Exec(query, category.Name, category.ID)
//	if err != nil {
//		log.Printf("Error updating category: %v", err)
//		return err
//	}
//
//	return nil
//}
//
//// DeleteCategoryService deletes a category by its ID
//func DeleteCategoryService(id int) error {
//	if database.DataBase == nil {
//		log.Println("Error: Database connection is nil")
//		return sql.ErrNoRows
//	}
//
//	query := "DELETE FROM Categories WHERE id = ?"
//	_, err := database.DataBase.Exec(query, id)
//	if err != nil {
//		log.Printf("Error deleting category: %v", err)
//		return err
//	}
//
//	return nil
//}
//
//// GetAllCategoriesService retrieves all categories
//func GetAllCategoriesService() ([]models.Category, error) {
//	if database.DataBase == nil {
//		log.Println("Error: Database connection is nil")
//		return nil, sql.ErrNoRows
//	}
//
//	rows, err := database.DataBase.Query("SELECT id, name FROM Categories")
//	if err != nil {
//		log.Printf("Error querying categories: %v", err)
//		return nil, err
//	}
//	defer rows.Close()
//
//	var categories []models.Category
//	for rows.Next() {
//		var category models.Category
//		if err := rows.Scan(&category.ID, &category.Name); err != nil {
//			log.Printf("Error scanning category: %v", err)
//			return nil, err
//		}
//		categories = append(categories, category)
//	}
//
//	return categories, nil
//}
