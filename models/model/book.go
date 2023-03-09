package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID           uint64         `json:"id"`
	Title        string         `json:"title"`
	BookCategory []BookCategory `json:"book_category"`
	BookAuthor   []BookAuthor   `json:"book_author"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
