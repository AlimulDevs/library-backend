package dto

import (
	"time"

	"gorm.io/gorm"
)

type BorrowBook struct {
	ID        uint64         `json:"id"`
	BorrowID  uint64         `json:"borrow_id"`
	BookID    uint64         `json:"book_id"`
	Book      Book           `json:"book"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
