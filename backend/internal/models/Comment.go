package models

import "time"

type Comment struct {
	ID          int       `json:"id"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"created_time"`
}
