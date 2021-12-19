package main

// gorm crud demo
import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
	"i-go/core/conf"
	"i-go/core/db/mysqldb"
)

var mysqlDB *gorm.DB

// // TableName 单独指定表名 设置User的表名为`profiles`
// func (User) TableName() string {
// 	return "profiles"
// }

type Animal struct {
	gorm.Model        // 直接将 base model 嵌入到自定义结构体中
	Name       string `gorm:"not null"`
	Age        int64
}

func main() {
	if err := conf.Load("../data/conf/config.yml"); err != nil {
		panic(err)
	}
	mysqldb.Init()
	mysqlDB = mysqldb.MySQL
	Create()
	Query()
	Update()
	Delete()
}

func Create() {
	err := mysqlDB.AutoMigrate(&Animal{})
	if err != nil {
		panic(err)
	}
	animal := Animal{
		Name: "cat",
		Age:  1,
	}
	err = mysqlDB.Model((*Animal)(nil)).Create(&animal).Error
	if err != nil {
		log.Println(err)
	}
	fmt.Println("id:", animal.ID) // 会通过 animal 对象返回id
	out := &Animal{}
	mysqlDB.Where("id = ? ", animal.ID).Find(out)
	fmt.Println(out)
}

func Delete() {
	// 普通删除 如果有deleteat字段就是软删除 没有就是物理删除
	mysqlDB.Where("id= ? ", 1).Delete(Animal{})
	// 强制物理删除 有deleteat字段也是物理删除
	mysqlDB.Where("id= ? ", 1).Unscoped().Delete(Animal{})
	// 普通查询-自定过滤掉软删除的记录
	var list = make([]Animal, 0)
	mysqlDB.Find(&list)
	fmt.Println("普通查询", list)
	// 特殊查询-会查询出软删除的记录
	mysqlDB.Unscoped().Find(&list)
	fmt.Println("Unscoped", list)
}

func Query() {
	// // 基本查询
	// var a Animal
	// mysqlDB.Where(Animal{Name: "cat"}).Find(&a).Limit(1)
	// fmt.Println(a)

	// limitOffset
	outs := make([]Animal, 0)
	mysqlDB.Order("id desc").Limit(3).Offset(1).Find(&outs)
	fmt.Printf("outs:%+v", outs)
	// count
	var c int64
	mysqlDB.Model((*Animal)(nil)).Count(&c)
	fmt.Println("c:", c)
	// groupHaving
	// groupHaving()
}

func Update() {
	var a = Animal{
		Model: gorm.Model{
			ID:        1, // 指定主键更新
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{
				Time:  time.Time{},
				Valid: false,
			},
		},
	}
	// 推荐：指定主键的多字段的更新，指定要更新的字段，*为全字段(包括 createAt等时间字段也会更新，所以都需要手动赋值)
	// createAt等时间字段，再不指定更新时也会自动更新，但是如果指定了要更新则需要手动指定值，不会自动生成了。
	mysqlDB.Model(&a).Select("Name", "Age").Updates(Animal{Name: "new_name", Age: 0})
	mysqlDB.Model(&a).Select("*").Updates(Animal{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Now(),
				Valid: false,
			},
		},
		Name: "dog",
		Age:  0,
	})

	// 推荐：指定更新条件的多字段的更新
	mysqlDB.Model(Animal{}).Where("name = ?", "cat").Updates(Animal{Name: "cat-new", Age: 18})

}

func groupHaving() {
	rows, err := mysqlDB.Table("animals").Select("COUNT(name) AS cn,age,sex").Where("age > ?", 10).Group("age,sex").
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
