package model

import (
	"time"
	//"gorm.io/gorm"
)

type (
	Customer struct {
		//gorm.Model
		ID         string `gorm:"type:varchar(256)"`
		Name       string `gorm:"type:varchar(256)"`
		Address    string `gorm:"type:varchar(256)"`
		Phone      string `gorm:"type:varchar(256)"`
		License_id string `gorm:"type:varchar(256)"`
		Active     bool   `gorm:"type:bool"`
		Email      string `gorm:"type:varchar(256)"`
		Password   string `gorm:"type:varchar(256)"`
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	ChangePasswordCondition struct {
		Id       string
		Password string
	}
)
