package models

import (
	"time"
)

type Category struct {
	ID              uint      `gorm:"column:id"`
	CategoryId      string    `gorm:"column:category_id"`
	Name            string    `gorm:"column:name"`
	Type            string    `gorm:"column:type"`
	IsSystemDefault bool      `gorm:"column:is_system_default"`
	UserId          string    `gorm:"column:user_id"`
	CreatedAt       time.Time `gorm:"column:created_at"`
}

func (Category) TableName() string {
	return "categories"
}
