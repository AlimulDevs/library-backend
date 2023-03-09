package model

import (
	"time"

	"gorm.io/gorm"
)

type BookAuthor struct {
	ID        uint64         `json:"id"`
	AuthorID  uint64         `json:"author_id"`
	BookID    uint64         `json:"book_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
