package models

import (
	"time"

	"gorm.io/gorm"
)

type OffersRegData struct {
	Id         uint64         `json:"id" db:"id" gorm:"primaryKey"`
	OfferId    uint64         `json:"offerId" db:"offer_id"`
	Price      float64        `json:"price" db:"price"`
	Percent    float64        `json:"percent" db:"percent"`
	PriceFinal float64        `json:"priceFinal" db:"price_final"`
	CreatedAt  time.Time      `json:"createdAt" db:"created_at"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" db:"deleted_at"`
}
