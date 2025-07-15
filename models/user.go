// models/user.go
package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `json:"-"` // hashed password
	CreatedAt time.Time `json:"created_at"`
}
