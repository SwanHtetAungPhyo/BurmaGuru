package models

import (
	"database/sql"
)

type Article struct {
	ID         int          `json:"id"`
	Title      string       `json:"title"`
	Content    string       `json:"content"`
	CategoryID int          `json:"category_id"`
	CountryID  int          `json:"country_id"`
	AuthorID   int          `json:"author_id"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}
