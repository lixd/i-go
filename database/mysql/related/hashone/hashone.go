package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"i-go/core/db/mysqldb"
)

// User 只能有一张信用卡 (CreditCard), CreditCardID 是外键
type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard `gorm:"foreignkey:uid;association_foreignkey:name"`
}

var mysqlDB = mysqldb.MySQL

func main() {
	mysqlDB.AutoMigrate(&CreditCard{})
	mysqlDB.AutoMigrate(&User{})
	var user User
	var card CreditCard
	mysqlDB.Model(&user).Related(&card, "CreditCard")
	//// SELECT * FROM credit_cards WHERE user_id = 123; // 123 is user's primary key
	// CreditCard 是 users 的字段，其含义是，获取 user 的 CreditCard 并填充至 card 变量
	// 如果字段名与 model 名相同，比如上面的例子，此时字段名可以省略不写，像这样：
	mysqlDB.Model(&user).Related(&card)
	fmt.Println("user:", user)
	fmt.Println("card:", card)
}
