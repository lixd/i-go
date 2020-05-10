package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"i-go/core/db/mysqldb"
)

var mysqlDB = mysqldb.MySQL

type User struct {
	gorm.Model
	Name        string `gorm:"type:varchar(20)"`
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
type Card struct {
	gorm.Model
	Attribute1      string
	Attribute2      string
	AssociatedCards []Card
}

/*type User struct {
	gorm.Model
	MemberNumber string
	CreditCards  []CreditCard `gorm:"foreignkey:UserMemberNumber;association_foreignkey:MemberNumber"`
}

type CreditCard struct {
	gorm.Model
	Number           string
	UserMemberNumber string
}*/

func main() {
	// create()
	bind()
	mysqlDB.AutoMigrate(&CreditCard{})
	mysqlDB.AutoMigrate(&User{})

	var cs []CreditCard
	xiaohan := &User{
		Model: gorm.Model{
			ID: 1,
		},
	}

	// 所有与xiaohan(id=1)相关联的CreditCard,找出来
	mysqlDB.Model(xiaohan).Association("CreditCards").Find(&cs)
	fmt.Printf("%#v\n", cs)
}

func bind() {
	// 创建一张卡片, 不做任何关联关系, ID自增
	mysqlDB.AutoMigrate(&Card{})
	mysqlDB.Create(&Card{})

	// 创建一张卡片, 同时关联一张卡片
	// 这里如果卡片存在, 直接关联, 如果不存在则会为你关联
	mysqlDB.Create(&Card{
		AssociatedCards: []Card{
			{
				Model: gorm.Model{ID: 2},
			},
		},
	})

	// 只关联两张卡片,将4&7两张卡片关联起来
	mysqlDB.Model(&Card{Model: gorm.Model{ID: 7}}).
		Association("AssociatedCards").
		Append(&Card{Model: gorm.Model{ID: 4}})

}

func create() {
	var u User
	mysqlDB.Create(&User{Name: "JKL"}).First(&u)
	mysqlDB.Create(&CreditCard{Number: "1", UserID: u.ID})
	mysqlDB.Create(&CreditCard{Number: "2", UserID: u.ID})
	mysqlDB.Create(&CreditCard{Number: "3", UserID: u.ID})
}
