package main

// gorm crud demo
import (
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
	Name string `gorm:"not null"`
	Age  int64
}

type T struct {
	ID   int64
	Name string
	Age  int
	Sex  string
}
type SoftDelete struct {
	gorm.Model
	Name string `sql:"not null"`
}

func main() {
	// Create()
	Query()
	// Update()
	// Delete()
	// softDelete()
}
func softDelete() {
	var sf1 = SoftDelete{Name: "first"}
	var sf2 = SoftDelete{Name: "second"}
	// mysqlDB.CreateTable(&sf1)
	mysqlDB.Create(&sf1)
	mysqlDB.Create(&sf2)
	var sfs = make([]SoftDelete, 0)
	mysqlDB.Find(&sfs)
	fmt.Println("普通查询", sfs)
	// 软删除
	mysqlDB.Where("Id = ?", 3).Delete(SoftDelete{})
	// 普通查询
	mysqlDB.Find(&sfs)
	fmt.Println("普通查询2", sfs)
	// 查询已删除的记录
	mysqlDB.Unscoped().Find(&sfs)
	fmt.Println("Unscoped", sfs)

	// 	物理删除
	mysqlDB.Where("Id = ?", 3).Unscoped().Delete(SoftDelete{})
	// 再次查询 真的删掉了
	mysqlDB.Unscoped().Find(&sfs)
	fmt.Println("Delete Unscoped", sfs)
}

func Delete() {
	var a Animal
	// 软删除 如果有deleteat字段就是软删除 没有就是物理删除
	mysqlDB.Where("id= ? ", 1).Delete(&a)
	// 物理删除 有deleteat字段也是物理删除
	mysqlDB.Where("id= ? ", 1).Unscoped().Delete(&a)
}

func Query() {
	// // 	子查询
	// subQuery()
	// limitOffset()
	groupHaving()
	/*	var a Animal
		mysqlDB.First(&a)
		fmt.Println(a)
		fmt.Printf("%#v\n", a)
		fmt.Printf("%+v\n", a)
		// 创建表时添加表后缀
		mysqlDB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})*/

}

func Update() {
	/*	// save
		var a1 Animal
		mysqlDB.Find(&a1)
		fmt.Println(a1)
		a1.Name = "NewName"
		mysqlDB.Save(&a1)*/
	// 	update
	var a1 = Animal{Age: 1, Name: "NewNameByUpdate"}
	// where条件必须写在update前面
	// mysqlDB.Model(&a1).Where("Id=?", 12).Update("name", "2NewNameByUpdate")

	// struct或map更新多个字段 当使用 struct 更新时，GORM只会更新那些非零值的字段
	// mysqlDB.Model(&a1).Where("Id=?", 12).Updates(Animal{Age: 1, Name: "xxx"})

	// Select或Omit 指定需要更新的字段
	// 只会更新age
	mysqlDB.Model(&a1).Where("Id=?", 12).Select("age").Updates(map[string]interface{}{"age": 3, "name": "zzz"})
	// 除了age都更新(只会更新updates中指定了的)
	mysqlDB.Model(&a1).Where("Id=?", 12).Omit("age").Updates(map[string]interface{}{"age": 3, "name": "ccc"})

}

func Create() {
	var id int64 = 20
	mysqlDB.CreateTable(&Animal{})
	var animal = Animal{Age: id}
	mysqlDB.Create(&animal)
	// INSERT INTO animals("age") values('99');
	// SELECT name from animals WHERE Id=111; // 返回主键为 111
	// animal.Name => 'galeone'
	out := &Animal{}
	outs := make([]Animal, 0)
	// mysqlDB.Find(out, "Id= ? ", id, )
	// mysqlDB.Where("Id= ? ", id).Find(out)
	mysqlDB.Where("Id!= ? ", id).Or("Id!= ? ", 1).Find(&outs)
	mysqlDB.Where("age = ? ", id).Attrs(animal).FirstOrCreate(&animal)
	// mysqlDB.Not("Id= ? ", id).Find(&outs)
	fmt.Println("name: ", out.Name)
	fmt.Println(outs)
}

func subQuery() {
	outs := make([]Animal, 0)
	mysqlDB.Where("age > ?", mysqlDB.Table("x_animals").Select("AVG(age)").SubQuery()).Find(&outs)
	fmt.Println(outs)
	var avg int64
	mysqlDB.Table("x_animals").Select("AVG(age)").Scan(&avg)
	fmt.Println("avg: ", avg)
}

func limitOffset() {
	outs2 := make([]Animal, 0)
	outs := make([]Animal, 0)
	var c int64
	mysqlDB.Offset(1).Order("age desc").Limit(3).Find(&outs).Order("age", true).Limit(-1).Find(&outs2).Count(&c)
	fmt.Println(c)
	fmt.Println(outs)
	fmt.Println(outs2)
}

func groupHaving() {
	rows, err := mysqlDB.Table("t").Select("COUNT(name) AS cn,age,sex").Where("age > ?", 10).Group("age,sex").
		Having("cn < ?", 2).Order("age").Rows()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// 分别scan各个值
	// for rows.Next() {
	// 	var cn int
	// 	var age int64
	// 	var sex string
	// 	rows.Scan(&cn, &age, &sex)
	// 	fmt.Println(cn, age, sex)
	// }
	// 一次性scan到model中
	var res Res
	for rows.Next() {
		mysqlDB.ScanRows(rows, &res)
		fmt.Printf("%#v\n", res)
	}
}

type Res struct {
	Cn  int
	Age int64
	Sex string
}
