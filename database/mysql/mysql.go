package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"i-go/core/db/mysqldb"
)

var mysqlDB = mysqldb.MySQL

// 使用指针
type User struct {
	gorm.Model
	Name     string `gorm:"default:'galeone',NOT NULL"`
	Age      int
	Password string `gorm:"NOT NULL"`
}

// TableName 单独指定表名 设置User的表名为`profiles`
func (User) TableName() string {
	return "profiles"
}

// BeforeCreate create之前可以对字段的值进行处理
// 比例将明文密码加密
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("password", encodePwd(user.Password))
	return nil
}
func encodePwd(pwd string) string {
	return pwd
}

type Animal struct {
	ID int64
	// Name string `gorm:"default:'galeone'"`
	Name sql.NullString `gorm:"not null"`
	Age  int64
}

func main() {
	var id int64 = 19
	// mysqlDB.CreateTable(&Animal{})
	var animal = Animal{Age: id}
	mysqlDB.Create(&animal)
	// INSERT INTO animals("age") values('99');
	// SELECT name from animals WHERE ID=111; // 返回主键为 111
	// animal.Name => 'galeone'
	out := &Animal{}
	// mysqlDB.Find(out, "ID= ? ", id, )
	mysqlDB.Where(out, "ID= ? ", id)
	fmt.Println("name: ", out.Name)
}
