package models

import "time"

type User struct {
	ID           uint
	UserId       string
	Email        string `gorm:"unique"`
	FirstName    string
	MidName      *string
	LastName     string
	PhoneNo      string
	Password     string
	CreatedAt    time.Time
	ImageProfile *string
	UpdatedAt    time.Time
}
