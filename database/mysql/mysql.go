package main

import (
	"i-go/core/db/mysqldb"
)

var mysqlDB = mysqldb.MySQL

type UserBasic struct {
	Id                int
	UnionId           string
	MpOpenId          string //小程序
	PubOpenId         string //公众号
	NickName          string `sql:",unique"`
	Avatar            string
	Sex               int
	Province          string
	City              string
	Country           string
	CountryCode       string
	PhoneNumber       string  `sql:",unique"`
	Description       string  //自我简述
	RegisterFrom      int     //注册来源
	Ip                string  //注册ip
	LoginIp           string  // 登陆ip
	RegisterTime      int64   // Seconds
	Status            int     `sql:"default:1"`
	Star              float64 `sql:"default:0"` //积分
	AlipayAccount     string  //支付宝账户
	WithdrawType      int     `sql:"default:1"` //提现账户 alipay
	ActiveShow        int     `sql:"default:0"`
	CommonwealVoteNum int64   //上一次参与的公益投票序号
	Note              string  //备注
}

func main() {

	//// 自动迁移模式
	//db.AutoMigrate(&Product{})

	// 创建
	//db.CreateTable(Product{})
	//db.Create(&Product{Code: "L1212", Price: 1000})

	//// 读取
	//var product Product
	////db.First(&product, 1)                   // 查询id为1的product
	////db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	//db.Find(&product, 1)
	//fmt.Println(product)
	//// 更新 - 更新product的price为2000
	//db.Model(&product).Update("Price", 2000)
	//
	//// 删除 - 删除product
	mysqlDB.Delete(UserBasic{}, "code = ?", "L1212")
}
