package model

import (
	"time"
)

type Flight struct {
	ID             string    `gorm:"type:varchar(256)"`
	Name           string    `gorm:"type:varchar(256)"`
	From           string    `gorm:"type:varchar(256)"`
	To             string    `gorm:"type:varchar(256)"`
	Date           time.Time `gorm:"type:timestamp"`
	Status         string    `gorm:"type:varchar(256)"`
	Available_slot int       `gorm:"type:integer"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
