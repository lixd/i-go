package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"i-go/core/conf"
	"i-go/core/db/mysqldb"
	"time"
)

var db *gorm.DB

func Init() {
	path := `D:\lillusory\projects\i-go\\conf\config.yml`
	conf.Load(path)
	mysqldb.Init()
	db = mysqldb.MySQL
	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})
}

func main() {
	Init()
	create()
	query()
	//update()
}

type User struct {
	//ID        uint       `gorm:"primary_key"`
	//CreatedAt time.Time  `gorm:"type:timestamp;"`
	//UpdatedAt time.Time  `gorm:"type:timestamp;"`
	//DeletedAt *time.Time `gorm:"type:timestamp;index:idx_del"`
	gorm.Model
	Name         string    `gorm:"type:varchar(20);unique_index:idx_name;NOT NULL;default:'guest'"`
	Phone        string    `gorm:"type:varchar(20);unique_index;NOT NULL"`
	Password     string    `gorm:"type:varchar(32);NOT NULL;default:'123456'"`
	LoginIp      string    `gorm:"type:varchar(64);NOT NULL"`                // 登陆ip
	LoginTime    time.Time `gorm:"type:timestamp;NOT NULL"`                  // 上次登陆时间
	RegisterIp   string    `gorm:"type:varchar(64);NOT NULL"`                // 注册ip
	RegisterTime time.Time `gorm:"type:timestamp;NOT NULL"`                  // time
	Status       int       `gorm:"type:tinyint UNSIGNED;NOT NULL;default:1"` // 状态
	Sex          int       `gorm:"type:tinyint UNSIGNED;NOT NULL"`
	Remark       string    `gorm:"type:varchar(64)"` // 备注
}

func create() {
	/*	user := User{
		Name:         "root",
		Phone:        "13452340416",
		Password:     "123456",
		LoginIp:      "127.0.0.1",
		LoginTime:    time.Now(),
		RegisterIp:   "192.168.1.1",
		RegisterTime: time.Now(),
		Status:       1,
		Sex:          1,
		Remark:       "first user",
	}	*/
	user := User{
		Name:         "admin",
		Phone:        "13452340417",
		Password:     "root",
		LoginIp:      "127.0.0.1",
		LoginTime:    time.Now(),
		RegisterIp:   "192.168.1.1",
		RegisterTime: time.Now(),
		Status:       1,
		Sex:          2,
		Remark:       "second user",
	}
	fmt.Println(db.NewRecord(&user))
	db.Create(&user)
	fmt.Println(db.NewRecord(&user))
}

func query() {
	//simple()
	//where()
	//structMap()
	//not()
	//or()
	//inlineCondition()
	//extraOptions()
	//order()
	groupHaving()
	//joins()
}

func update() {
	var user User
	// 更新单个属性，如果它有变化
	db.Model(&user).Where("name = ?", "root").Update("status", 2)
	//	 使用 SQL 表达式更新
	db.Model(&user).Where("name = ?", "root").Update("status", gorm.Expr("status*2"))
}

const (
	Layout = "2006-01-02 15:04:05"
)

func simple() {
	var users = make([]User, 0)
	db.Find(&users)
	fmt.Println(users[0].LoginTime.Format(Layout))
}

func where() {
	var user User
	db.Where("name = ?", "root").Find(&user)
	fmt.Println(user)
	// Find 传入单个结构体对象则会返回一个值 传入切片则返回多个值
	//var userTime User
	var users = make([]User, 0)
	//db.Where("register_time < ?", time.Now()).Find(&userTime)
	db.Where("register_time < ?", time.Now()).Find(&users)
	fmt.Println(users)
	// BETWEEN  AND
	var between = make([]User, 0)
	db.Where("register_time BETWEEN ? AND ?", time.Now().Add(-time.Minute*10), time.Now()).Find(&between)
	fmt.Println(between)
}
func structMap() {
	// Struct
	var userStruct User
	db.Where(&User{Name: "root"}).Find(&userStruct)
	fmt.Println(userStruct)
	//	Map
	var userMap User
	db.Where(map[string]interface{}{"name": "root"}).Find(&userMap)
	fmt.Println(userMap)
}
func not() {
	var user User
	db.Not(&User{Name: "root"}).Find(&user)
	fmt.Println(user)
}
func or() {
	var user User
	db.Where("name = ?", "root").Or("name = ?", "admin").Find(&user)
	fmt.Println(user)
}
func inlineCondition() {
	var user User
	db.Find(&user, "name = ?", "root")
	fmt.Println(user)
}
func extraOptions() {
	var user User
	// 为查询 SQL 添加额外的 SQL 操作
	db.Set("gorm:query_option", "FOR UPDATE").First(&user, 1)
	//// SELECT * FROM users WHERE id = 10 FOR UPDATE;
}
func order() {
	var users = make([]User, 0)
	db.Order("phone desc, name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;
	fmt.Println(users)
	// 多字段排序
	db.Order("phone desc").Order("name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;
	fmt.Println(users)
}
func groupHaving() {
	rows, err := db.Table("users").Select("count(id) as total").Group("status").Having("total > ?", 1).Rows()
	if err != nil {
		fmt.Println(err)
		return
	}
	var total int
	for rows.Next() {
		rows.Scan(&total)
		fmt.Println(total)
	}
}

/*func joins() {
	rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
	for rows.Next() {
		...
	}
}*/

func delete() {
}
