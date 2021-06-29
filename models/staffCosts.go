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
	Project1Name     string         `json:"project1Name" db:"project1_name"`
	Percentage1      float64        `json:"per1" db:"per1"`
	Project2         string         `json:"project2" db:"project2"`
	Project2Name     string         `json:"project2Name" db:"project2_name"`
	Percentage2      float64        `json:"per2" db:"per2"`
	Project3         string         `json:"project3" db:"project3"`
	Project3Name     string         `json:"project3Name" db:"project3_name"`
	Percentage3      float64        `json:"per3" db:"per3"`
	Project4         string         `json:"project4" db:"project4"`
	Project4Name     string         `json:"project4Name" db:"project4_name"`
	Percentage4      float64        `json:"per4" db:"per4"`
	PayDate          time.Time      `json:"payDate" db:"pay_date"`
	CreatedAt        time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time      `json:"updatedAt" db:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt" db:"deleted_at"`
}

type StaffProject struct {
	Id         uint64         `json:"id" db:"id" gorm:"primaryKey"`
	IdStaff    uint64         `json:"idStaff" db:"id_staff"`
	Project    string         `json:"project" db:"project"`
	Amount     float64        `json:"amount" db:"amount"`
	Percentage float64        `json:"percentaje" db:"percentaje"`
	Total      float64        `json:"total" db:"total"`
	CreatedAt  time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time      `json:"updatedAt" db:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" db:"deleted_at"`
}
