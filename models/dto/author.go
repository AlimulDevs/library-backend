package dto

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID         uint64         `json:"id"`
	Name       string         `json:"name"`
	BookAuthor []BookAuthor   `json:"book_author"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
