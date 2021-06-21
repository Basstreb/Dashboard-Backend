package models

import (
	"time"

	"gorm.io/gorm"
)

type IvaAmountModel struct {
	Month  int     `json:"month" db:"month"`
	Amount float64 `json:"amount" db:"amount"`
}

type IvaModel struct {
	Name string    `json:"name" db:"name"`
	Iva  float64   `json:"iva" db:"iva"`
	Date time.Time `json:"date" db:"date"`
}

type IvaAccumulated struct {
	Id        uint64         `json:"id" db:"id" gorm:"primaryKey"`
	Amount    float64        `json:"amount" db:"amount"`
	CreatedAt time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time      `json:"updatedAt" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" db:"deleted_at"`
}

type IvaPaid struct {
	Id     uint64  `json:"id" db:"id" gorm:"primaryKey"`
	Amount float64 `json:"amount" db:"amount"`
}
