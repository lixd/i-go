package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"i-go/database/postgresql/constant"
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
	// 1 user 1 &{1 en}
	// 2 user 2 &{2 ru}
}

func initTable() {
	qs := []string{
		"CREATE TEMP TABLE products (id int, name text)",
		"CREATE  TEMP TABLE categories (id int, description text,product_id int)",
		"INSERT INTO products VALUES (1, 'product 1'), (2, 'product 2')",
		"INSERT INTO categories VALUES (1, 'meat',1), (2, 'fruit',2)",
	}
	for _, q := range qs {
		_, err := pgDB2.Exec(q)
		if err != nil {
			panic(err)
		}
	}
}

func connect() {
	pgDB2 = pg.Connect(&pg.Options{
		User:     constant.UserName,
		Addr:     constant.Addr,
		Password: constant.Password,
		Database: constant.Database,
	})
}

func init() {
	connect()
	initTable()
}
