package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"i-go/core/db/mysqldb"
)

var mysqlDB = mysqldb.MySQL

/*// `User`属于`Profile`, `ProfileID`为外键
// 默认外键名为 表名+主键名(tbl_name+PK)
// 默认关联的外键为PK
type User struct {
	gorm.Model
	Profile   Profile
	ProfileID int
}

type Profile struct {
	gorm.Model
	Name string
}
*/

/*// 指定外键
// 外键列名就不用满足上面的规范了表名+主键名(tbl_name+PK)
// 关联的外键还是默认为主键
type User struct {
	gorm.Model
	Profile      Profile `gorm:"ForeignKey:ProfileRefer"` // 使用ProfileRefer作为外键
	ProfileRefer int
}

type Profile struct {
	gorm.Model
	Name string
}*/

// 指定外键和关联外键
// 则二者都不需要满足默认规范了
type Profile struct {
	gorm.Model
	Refer string
	Name  string
}

type User struct {
	gorm.Model
	Profile   Profile `gorm:"ForeignKey:ProfileID;AssociationForeignKey:Refer"`
	ProfileID int
}

func main() {
	// mysqlDB.CreateTable(&User{})
	// mysqlDB.CreateTable(&Profile{})
	mysqlDB.AutoMigrate(&User{})
	mysqlDB.AutoMigrate(&Profile{})
	var p = Profile{Name: "p1"}
	mysqlDB.Create(&User{Profile: p})
	mysqlDB.Create(&p)
	var User = User{}
	mysqlDB.Model(&User).Related(&User.Profile, "ProfileID").Create(User)
	fmt.Println("user: ", User)

}
