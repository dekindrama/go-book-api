package models

import (
	"time"

	"gorm.io/gorm"
)

type BookModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Author    string         `gorm:"not null" json:"author"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (BookModel) TableName() string {
	return "books"
}
