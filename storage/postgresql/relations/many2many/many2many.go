/*
relation-has one
*/
package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"i-go/storage/postgresql/constant"
	"i-go/storage/postgresql/util"
)

type Category struct {
	Id          int
	Description string
}

type Product struct {
	Id       int
	Products []Product `pg:"many2many:product_to_products,joinFK:category_id"`
}

type ProductToProducts struct {
	CategoryId int
	ProductId  int
}

var pgDB2 *pg.DB

func main() {
	defer pgDB2.Close()
	var Product Product
	err := pgDB2.Model(&Product).
		Column("product.*").
		Relation("Products").
		First()
	if err != nil {
		panic(err)
	}

	fmt.Println(Product.Id, Product.Products)
	// fmt.Println(Product[1].Id, Product[1].Products)
	// Output: 2 results
	// 1 product 1 [0xc000137200 0xc000137240]
	// 2 product 1 [0xc000137280 0xc0001372c0]
}

func initTable() {
	models := []interface{}{
		(*Category)(nil),
		(*Product)(nil),
		(*ProductToProducts)(nil),
	}
	for _, v := range models {
		err := pgDB2.CreateTable(v, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
			Temp:          true,
		})
		util.HandError("CreateTable Category", err)
	}
	orm.RegisterTable((*ProductToProducts)(nil))
}

func initData() {
	ps := []interface{}{
		&Product{1, nil},
		&Product{2, nil},
		&Product{3, nil},
		&Category{1, "Category 1"},
		&Category{2, "Category 1"},
		&Category{3, "Category 1"},
		&ProductToProducts{1, 1},
		&ProductToProducts{2, 2},
		&ProductToProducts{3, 3},
	}
	for _, v := range ps {
		err := pgDB2.Insert(v)
		if err != nil {
			panic(err)
		}
	}
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
