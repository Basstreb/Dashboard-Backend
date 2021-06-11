package models

import (
	"time"

	"gorm.io/gorm"
)

type TypoCosts struct {
	Id        uint64         `json:"id" db:"id" gorm:"primaryKey"`
	Typo      string         `json:"typo" db:"typo"`
	Iva       uint64         `json:"iva" db:"iva"`
	CreatedAt time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time      `json:"updatedAt" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" db:"deleted_at"`
}
