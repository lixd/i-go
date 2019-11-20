package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Order struct {
	Id    int
	Items []Item `pg:"many2many:order_to_items"`
}

type Item struct {
	Id int
}

type OrderToItem struct {
	OrderId int
	ItemId  int
}

func createManyToManyTables(db *pg.DB) error {
	models := []interface{}{
		(*Order)(nil),
		(*Item)(nil),
		(*OrderToItem)(nil),
	}
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
