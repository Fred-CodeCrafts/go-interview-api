// models/question.go
package models

import "time"

type Question struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Topic     string    `json:"topic"`
	Question  string    `json:"question"`
	AuthorID  uint      `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}
