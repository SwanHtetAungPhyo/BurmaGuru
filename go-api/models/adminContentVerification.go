package models

import "time"


type AdminContentVerification struct {
    ID         int       `json:"id"`
    ContentID  int       `json:"content_id"`
    ContentType string   `json:"content_type"` // 'article' or 'guide'
    AuthorID   int       `json:"author_id"`
    Status     string    `json:"status"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}