package belongto

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"i-go/core/db/mysqldb"
)

var mysqlDB = mysqldb.MySQL

/*
// 默认外键名为 表名+主键名(tbl_name+PK)
// 默认关联的外键为PK
type User struct {
  gorm.Model
  Name string
}

// `Profile` 属于 `User`， 外键是`UserId`
type Profile struct {
  gorm.Model
  UserId int
  User   User
  Name   string
}
*/

/*// 指定外键
// 外键列名就不用满足上面的规范了表名+主键名(tbl_name+PK)
// 关联的外键还是默认为主键
type User struct {
  gorm.Model
  Name string
}

type Profile struct {
  gorm.Model
  Name      string
  User      User `gorm:"foreignkey:UserRefer"` // 将 UserRefer 作为外键
  UserRefer uint
}*/

// 指定外键和关联外键
// 则二者都不需要满足默认规范了
/*type Profile struct {
	gorm.Model
	Refer string
	Name  string
}

type User struct {
	gorm.Model
	Profile   Profile `gorm:"ForeignKey:ProfileID;AssociationForeignKey:Refer"`
	ProfileID int
}*/
type User struct {
	gorm.Model
	Refer string
	Name  string
}

type Profile struct {
	gorm.Model
	Name string
	//User      User `gorm:"association_foreignkey:Refer"` // 将 Refer 作为关联外键
	User      User `gorm:"ForeignKey:UserRefer;AssociationForeignKey:Refer"` // 将 Refer 作为关联外键
	UserRefer string
}

func main() {
	// mysqlDB.CreateTable(&User{})
	// mysqlDB.CreateTable(&Profile{})
	mysqlDB.AutoMigrate(&User{})
	mysqlDB.AutoMigrate(&Profile{})
	var p = Profile{Name: "p1"}
	mysqlDB.Create(&User{})
	mysqlDB.Create(&p)
	var User = User{}
	mysqlDB.Model(&User).Related(&User, "ProfileID").Create(User)
	fmt.Println("user: ", User)

}
