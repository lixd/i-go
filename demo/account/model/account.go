package model

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	UserId uint    `gorm:"type:int;NOT NULL"`
	Amount float64 `gorm:"type:decimal(10,6)"`
}
