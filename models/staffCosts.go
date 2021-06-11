package models

import (
	"time"

	"gorm.io/gorm"
)

type StaffCosts struct {
	Id               uint64         `json:"id" db:"id" gorm:"primaryKey"`
	StaffName        string         `json:"staffName" db:"staff_name"`
	Amount           float64        `json:"amount" db:"amount"`
	Cost             float64        `json:"cost" db:"cost"`
	SocialInsurances float64        `json:"socialInsurances" db:"social_insurances"`
	Project1         string         `json:"project1" db:"project1"`
	Project2         string         `json:"project2" db:"project2"`
	Project3         string         `json:"project3" db:"project3"`
	Project4         string         `json:"project4" db:"project4"`
	PayDate          time.Time      `json:"payDate" db:"pay_date"`
	CreatedAt        time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time      `json:"updatedAt" db:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt" db:"deleted_at"`
}
