package model

import (
	"time"
)

type Booking struct {
	ID          string    `gorm:"type:varchar(256)"`
	Customer_id string    `gorm:"type:varchar(256)"`
	Flight_id   string    `gorm:"type:varchar(256)"`
	Code        string    `gorm:"type:varchar(256)"`
	Status      string    `gorm:"type:varchar(256)"`
	Booked_date time.Time `gorm:"type:timestamp"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
