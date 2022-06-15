package models

import (
	"gorm.io/gorm"
)

type UserData struct {
	Email    string `json:"email" db:"email"`
	Password []byte `json:"-" db:"password"`
	Name     string `json:"name" db:"name"`
	gorm.Model
}
