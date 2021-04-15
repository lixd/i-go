/*
relation-belongs to
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
	ProductId   int
}

// Product has one Category.
type Product struct {
	Id       int
	Name     string
	Category *Category
}

var pgDB2 *pg.DB

func main() {
	defer pgDB2.Close()

	var Products []Product
	err := pgDB2.Model(&Products).
		Column("product.*").
		Relation("Category.product_id").
		Select()
	if err != nil {
		panic(err)
	}

	fmt.Println(len(Products), "results")
	fmt.Println(Products[0].Id, Products[0].Name, Products[0].Category)
	fmt.Println(Products[1].Id, Products[1].Name, Products[1].Category)
	// Output: 2 results
	// 1 product 1 &{0  1}
	// 2 product 1 &{0  2}
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
	c1 := &Category{1, "category 1", 1}
	c2 := &Category{2, "category 2", 2}
	_, _ = pgDB2.Model(c1, c2).Insert()
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
