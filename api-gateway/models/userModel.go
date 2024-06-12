package models

import "time"

type User struct {
	UserID    int       `json:"user_id" gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
}

func (u *User) TableName() string {
	return "user"
}

type RegisterUser struct {
	UserID    int       `json:"user_id" gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
}

func (u *RegisterUser) TableName() string {
	return "user"
}
