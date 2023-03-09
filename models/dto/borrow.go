package dto

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	ID         uint64         `json:"id"`
	BorrowerID uint64         `json:"borrower_id"`
	Date       time.Time      `json:"date"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
