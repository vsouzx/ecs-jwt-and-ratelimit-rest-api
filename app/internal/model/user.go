package model

import "time"

type User struct {
	Identifier uint      `gorm:"primaryKey" json:"identifier"`
	Name       string    `gorm:"size:120;not null" json:"name"`
	Email      string    `gorm:"size:120;uniqueIndex;not null" json:"email"`
	Password   string    `gorm:"size:255;not null" json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
