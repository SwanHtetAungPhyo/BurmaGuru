package services

import (
	"database/sql"
	"log"
	"time"

	"github.com/swanhtetaungphyo/burmaguru/database"
	"github.com/swanhtetaungphyo/burmaguru/models"
)

// CreateArticleService creates a new article
func CreateArticleService(article models.Article) (models.Article, error) {
	if database.DataBase == nil {
		log.Println("Error: Database connection is nil")
		return models.Article{}, sql.ErrNoRows
	}

	query := `
		INSERT INTO Articles (title, content, category_id, country_id, author_id, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	// Create sql.NullTime from time.Time
	article.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	article.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	result, err := database.DataBase.Exec(query, article.Title, article.Content, article.CategoryID, article.CountryID, article.AuthorID, article.CreatedAt, article.UpdatedAt)
	if err != nil {
		log.Printf("Error inserting article: %v", err)
		return models.Article{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert id: %v", err)
		return models.Article{}, err
	}
	article.ID = int(id)

	return article, nil
}

// GetArticleByIdService retrieves an article by its ID
func GetArticleByIdService(id int) (models.Article, error) {
	if database.DataBase == nil {
		log.Println("Error: Database connection is nil")
		return models.Article{}, sql.ErrNoRows
	}

	var article models.Article
	query := "SELECT id, title, content, category_id, country_id, author_id, created_at, updated_at FROM Articles WHERE id = ?"
	err := database.DataBase.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content, &article.CategoryID, &article.CountryID, &article.AuthorID, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		log.Printf("Error retrieving article: %v", err)
		return models.Article{}, err
	}

	return article, nil
}

// UpdateArticleService updates an existing article
func UpdateArticleService(article models.Article) error {
	if database.DataBase == nil {
		log.Println("Error: Database connection is nil")
		return sql.ErrNoRows
	}

	article.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true} // Set UpdatedAt
	query := `
		UPDATE Articles 
		SET title = ?, content = ?, category_id = ?, country_id = ?, author_id = ?, updated_at = ? 
		WHERE id = ?`
	_, err := database.DataBase.Exec(query, article.Title, article.Content, article.CategoryID, article.CountryID, article.AuthorID, article.UpdatedAt, article.ID)
	if err != nil {
		log.Printf("Error updating article: %v", err)
		return err
	}

	return nil
}

// DeleteArticleService deletes an article by its ID
func DeleteArticleService(id int) error {
	if database.DataBase == nil {
		log.Println("Error: Database connection is nil")
		return sql.ErrNoRows
	}

	query := "DELETE FROM Articles WHERE id = ?"
	_, err := database.DataBase.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting article: %v", err)
		return err
	}

	return nil
}

func GetAllArticlesService() ([]models.Article, error) {
	if database.DataBase == nil {
		log.Println("Error: Database connection is nil")
		return nil, sql.ErrNoRows
	}

	rows, err := database.DataBase.Query("SELECT id, title, content, category_id, country_id, author_id, created_at, updated_at FROM Articles")
	if err != nil {
		log.Printf("Error querying articles: %v", err)
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.CategoryID, &article.CountryID, &article.AuthorID, &article.CreatedAt, &article.UpdatedAt); err != nil {
			log.Printf("Error scanning article: %v", err)
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
