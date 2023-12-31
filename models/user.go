// models/user.go

package models

import "time"

// User is a struct representing user data in the application.
type User struct {
	ID        string    `gorm:"primaryKey;not null" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null" json:"username"`
	Email     string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for the User model.
func (User) TableName() string {
	return "users"
}
