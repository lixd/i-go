package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `gorm:"type:varchar(16);"`
	Pwd        string `gorm:"type:varchar(16);NOT NULL"`
	Phone      string `gorm:"type:varchar(20);NOT NULL"`
	Age        uint   `gorm:"type:tinyint;"`
	RegisterIP string `gorm:"type:varchar(64);"`
	LoginIP    string `gorm:"type:varchar(64);"`
}
