package models

import (
	"time"

	"gorm.io/gorm"
)

type OffersData struct {
	Id           uint64         `json:"id" db:"id" gorm:"primaryKey"`
	ClientId     uint64         `json:"clientId" db:"client_id"`
	CompanyName  string         `json:"companyName" db:"company_name"`
	Offer        string         `json:"offer" db:"offer"`
	OfferName    string         `json:"offerName" db:"offer_name"`
	DecisionDate time.Time      `json:"decisionDate" db:"decision_date"`
	Code         string         `json:"code" db:"code"`
	Status       string         `json:"status" db:"status"`
	Percent      float64        `json:"percent" db:"percent"`
	Price        float64        `json:"price" db:"price"`
	PriceIva     float64        `json:"priceIva" db:"price_iva"`
	CreatedAt    time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time      `json:"updatedAt" db:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" db:"deleted_at"`
}

type MonthPriceOffer struct {
	Month int     `json:"month" db:"month"`
	Price float64 `json:"price" db:"price"`
}

type OfferRent struct {
	Id         uint64  `json:"id" db:"id"`
	OfferName  string  `json:"offerName" db:"offer_name"`
	Total      float64 `json:"total" db:"total"`
	CommonCost float64 `json:"commonCost" db:"common_cost"`
	StaffCost  float64 `json:"staffCost" db:"staff_cost"`
}
