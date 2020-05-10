package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"i-go/core/db/mysqldb"
)

var mysqlDB = mysqldb.MySQL

type Student struct {
	gorm.Model
	// PhoneNumber string `gorm:"UNIQUE;NOT NULL"` // 唯一、非空约束
	// 非空约束列推荐使用sql.NullXXX类型
	PhoneNumber sql.NullString `gorm:"type:varchar(20);UNIQUE;NOT NULL"`                      // 唯一约束和非空约束
	Name        string         `gorm:"column:newName;type:varchar(20);DEFAULT:'defaultName'"` // 列名类型默认值
	Sex         string         `gorm:"type:ENUM('F','M')"`                                    // 指定类型枚举
	Class       int            `gorm:"DEFAULT:2"`                                             // 默认值
	IgnoreMe    int            `gorm:"-"`                                                     // 忽略本字段
}

func main() {
	mysqlDB.CreateTable(Student{})
	mysqlDB.Create(&Student{Sex: "F"})
}
