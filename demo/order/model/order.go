package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId uint    `gorm:"type:int;NOT NULL"`
	Amount float64 `gorm:"type:decimal(10,6)"`
}
