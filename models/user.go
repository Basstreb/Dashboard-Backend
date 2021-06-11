package models

import (
	"time"

	"gorm.io/gorm"
)

type UserData struct {
	Id        uint64         `json:"id" db:"id"`
	Email     string         `json:"email" db:"email"`
	Password  []byte         `json:"-" db:"password"`
	Name      string         `json:"name" db:"name"`
	CreatedAt time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time      `json:"updatedAt" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" db:"deleted_at"`
}
