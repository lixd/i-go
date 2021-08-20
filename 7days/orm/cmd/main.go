package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	orm2 "i-go/7days/orm"
)

// func main() {
// 	db, _ := sql.Open("sqlite3", "17x.db")
// 	defer db.Close()
// 	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
// 	_, _ = db.Exec("CREATE TABLE User(Name text);")
// 	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
// 	if err != nil {
// 		log.Println("exec err: ", err)
// 		return
// 	}
// 	affected, _ := result.RowsAffected()
// 	log.Println("写入记录条数: ", affected)
// 	rows, err := db.Query("SELECT Name FROM User LIMIT 2")
// 	if err != nil {
// 		log.Println("query err: ", err)
// 		return
// 	}
// 	defer rows.Close()
// 	var name string
// 	for rows.Next() {
// 		if err := rows.Scan(&name); err == nil {
// 			log.Println(name)
// 		}
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Println("rows err: ", err)
// 	}
// }

func main() {
	engine, _ := orm2.NewEngine("sqlite3", "17x.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
