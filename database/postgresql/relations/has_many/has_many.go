/*
relation-has one
*/
package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"i-go/database/postgresql/constant"
	"i-go/database/postgresql/util"
)

type Category struct {
	Id          int
	Description string
	ProductId   int
}

// Product has many Category.
type Product struct {
	Id         int
	Name       string
	Categories []*Category
}

var pgDB2 *pg.DB

func main() {
	defer pgDB2.Close()

	var Product []Product
	err := pgDB2.Model(&Product).
		Column("product.*").
		Relation("Categories").
		Select()
	if err != nil {
		panic(err)
	}

	fmt.Println(len(Product), "results")
	fmt.Println(Product[0].Id, Product[0].Name, Product[0].Categories)
	fmt.Println(Product[1].Id, Product[1].Name, Product[1].Categories)
	// Output: 2 results
	// 1 product 1 [0xc000137200 0xc000137240]
	// 2 product 1 [0xc000137280 0xc0001372c0]
}

func initTable() {
	models := []interface{}{
		(*Product)(nil),
		(*Category)(nil),
	}
	for _, v := range models {
		err := pgDB2.CreateTable(v, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
			Temp:          true,
		})
		util.HandError("CreateTable Category", err)
	}
}

func initData() {
	p1 := &Product{1, "product 1", nil}
	p2 := &Product{2, "product 1", nil}
	_, _ = pgDB2.Model(p1, p2).Insert()
	c1 := &Category{1, "Product 1", 1}
	c2 := &Category{2, "Product 2", 1}
	c3 := &Category{3, "Product 3", 2}
	c4 := &Category{4, "Product 4", 2}
	_, _ = pgDB2.Model(c1, c2, c3, c4).Insert()
}

func initConn() {
	pgDB2 = pg.Connect(&pg.Options{
		User:     constant.UserName,
		Addr:     constant.Addr,
		Password: constant.Password,
		Database: constant.Database,
	})
}

func init() {
	initConn()
	initTable()
	initData()
}
