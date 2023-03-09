package dto

import (
	"time"

	"gorm.io/gorm"
)

type Borrower struct {
	ID         uint64         `json:"id"`
	Name       string         `json:"name"`
	NoIdentity string         `json:"no_identity" gorm:"uniqueIndex"`
	Address    string         `json:"address"`
	NoHp       string         `json:"no_hp" gorm:"uniqueIndex"`
	Borrow     []Borrow       `json:"borrow"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
