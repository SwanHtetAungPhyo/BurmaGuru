package models

import "time"

type Connection struct {
    ID           int       `json:"id"`
    UserID       int       `json:"user_id"`
    TargetUserID int       `json:"target_user_id"`
    Status       string    `json:"status"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
