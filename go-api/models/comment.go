package models

import "time"

type Comment struct {
    ID        int       `json:"id"`
    Content   string    `json:"content"`
    AuthorID  int       `json:"author_id"`
    ArticleID int       `json:"article_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
