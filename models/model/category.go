package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           uint64         `json:"id"`
	Name         string         `json:"name"`
	BookCategory []BookCategory `json:"book_category"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
