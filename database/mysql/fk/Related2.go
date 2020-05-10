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
	create()
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

func create() {
	var u User
	mysqlDB.Create(&User{Name: "JKL"}).First(&u)
	mysqlDB.Create(&CreditCard{Number: "1", UserID: u.ID})
	mysqlDB.Create(&CreditCard{Number: "2", UserID: u.ID})
	mysqlDB.Create(&CreditCard{Number: "3", UserID: u.ID})
}
