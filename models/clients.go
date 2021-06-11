package models

import (
	"time"

	"gorm.io/gorm"
)

type ClientData struct {
	Id            uint64         `json:"id" db:"id" gorm:"primaryKey"`
	CompanyName   string         `json:"companyName" db:"company_name"`
	ManagerName   string         `json:"managerName" db:"manager_name"`
	ManagerCharge string         `json:"managerCharge" db:"manager_charge"`
	PhoneNumber   string         `json:"phoneNumber" db:"phone_number"`
	Cif           string         `json:"cif" db:"cif"`
	Direction     string         `json:"direction" db:"direction"`
	CreatedAt     time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time      `json:"updatedAt" db:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt" db:"deleted_at"`
}
