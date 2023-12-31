package models

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;not null" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null" json:"username"`
	Email     string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
