package models

import (
	"time"

	"gorm.io/gorm"
)

type CommonCosts struct {
	Id              uint64         `json:"id" db:"id" gorm:"primaryKey"`
	CommonCostsName string         `json:"commonCostsName" db:"common_costs_name"`
	OfferId         uint64         `json:"offerId" db:"offer_id"`
	OfferName       string         `json:"offerName" db:"offer_name"`
	Amount          float64        `json:"amount" db:"amount"`
	AmountW         float64        `json:"amountW" db:"amount_w"`
	CostDate        time.Time      `json:"costDate" db:"cost_date"`
	CreatedAt       time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time      `json:"updatedAt" db:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deletedAt" db:"deleted_at"`
}

type MonthPriceCost struct {
	Month  int     `json:"month" db:"month"`
	Amount float64 `json:"amount" db:"amount"`
}
