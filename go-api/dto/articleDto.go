package dto

import "time"

type ArticleDto struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CategoryID int       `json:"category_id"`
	CountryID  int       `json:"country_id"`
	AuthorID   int       `json:"author_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
