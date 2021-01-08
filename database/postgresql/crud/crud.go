package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"i-go/core/conf"
	"i-go/core/db/pgsqldb"
	"i-go/database/postgresql/util"
)

var db *pg.DB

type Book struct {
	AuthorId string
	Id       int
	Title    string
	Text     string
	Author   *Author
}

type Author struct {
	Id     int
	Name   string
	BookId int
}

type Params struct {
	X int
	Y int
}

func (p *Params) Sum() int {
	return p.X + p.Y
}

func init() {
	initConn()
	initTable()
}

func initConn() {
	db = pgsqldb.PostgresDB
}

func Init(path string) {
	err := conf.Load(path)
	if err != nil {
		panic(err)
	}
}

func main() {
	Init("conf/config.json")

	defer pgsqldb.Release()
	ModelTest()
	// ExampleInsert2()
	// ExampleSelect()
	// ExampleJoin()
	//
	// ExampleColumn()
	//ExampleInsert()
	//
	// ExampleUpdate()
	// ExampleDelete()
	// ExampleCustomQuery()
	// ExampleQuery()
	//ExampleQuery2()
}

type User struct {
	Id   int    `pg:"pk"`
	Name string `pg:"notnull"`
}

func ModelTest() {
	err := db.CreateTable(User{}, &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	})
	if err != nil {
		panic(err)
	}
}

func ExampleQuery2() {
	var books []Book
	b := new(Book)
	db.Model(b).Where("id = ?", 1)
	books = append(books, *b)
	err := db.Model(&books).
		Column("book.*").
		Relation("Author").
		Select()
	if err != nil {
		panic(err)
	}
	for _, v := range books {
		fmt.Println(v)
	}
}

func ExampleCustomQuery() {
	// Create index for the table:
	customCreateIndex()

	customQuery()
}

func customQuery() {
	var count int
	_, err := db.Model((*Book)(nil)).QueryOne(pg.Scan(&count), `
	    SELECT count(*)
	    FROM ?TableName AS ?TableAlias
	`)
	util.HandError("custom", err)
}

func customCreateIndex() {
	_, err := db.Model((*Book)(nil)).Exec(`
	    CREATE INDEX CONCURRENTLY books_author_id_idx
	    ON ?TableName (author_id)
	`)
	util.HandError("custom", err)
}

func ExampleDelete() {
	// Delete book by primary key:
	deleteByPkey()
	// The same but more explicitly:
	deleteByPkey2()

	// Delete book by title:
	deleteByParam()

	// Delete multiple books:
	deleteMultiple()
}

func deleteMultiple() {
	books := []Book{} // slice of books with ids
	books = append(books, Book{AuthorId: "12345", Id: 4, Text: "insert test", Title: "insert"})
	books = append(books, Book{AuthorId: "12345", Id: 5, Text: "insert test 2", Title: "insert 2"})

	res, err := db.Model(&books).WherePK().Delete()
	util.HandError("deleteMultiple", err)
	fmt.Println(res)
	// DELETE FROM "books" WHERE id IN (1, 2, 3)
}

func deleteByParam() {
	b := &Book{AuthorId: "12345", Id: 4, Text: "insert test", Title: "insert"}
	res, err := db.Model(b).Where("title = ?title").Delete()
	util.HandError("delete", err)
	fmt.Println(res)
	// DELETE FROM "books" WHERE title = 'my title'
}

func deleteByPkey2() {
	b := &Book{AuthorId: "12345", Id: 4, Text: "insert test", Title: "insert"}
	res, err := db.Model(b).WherePK().Delete()
	util.HandError("delete", err)
	fmt.Println(res)
	// DELETE FROM "books" WHERE id = 1
}

func deleteByPkey() {
	b := &Book{AuthorId: "12345", Id: 4, Text: "insert test", Title: "insert"}
	err := db.Delete(b)
	util.HandError("delete", err)
	// DELETE FROM "books" WHERE id = 1
}

func ExampleInsert() {
	// Insert new book returning primary keys:
	insertReturningDefault()

	// Insert new book returning all columns:
	insertReturningAll()

	// Insert multiple books with single query:
	insertMultiple()

	// Select existing book by name or create new book:

	insertExisting()
	// Insert new book or update existing one:

	insertCreate()

}

func insertCreate() {
	author := &Author{Id: 123, Name: "author1"}
	_, err := db.Model(author).
		OnConflict("(id) DO UPDATE").
		Set("name = EXCLUDED.name").
		Insert()
	util.HandError("insert", err)
	fmt.Println(author)

	// INSERT INTO "books" ("id", "title") VALUES (100, 'my title')
	// ON CONFLICT (id) DO UPDATE SET title = 'title version #1'
}
func insertExisting() {
	b := &Book{AuthorId: "12345", Id: 4, Text: "insert test", Title: "insert"}
	_, err := db.Model(b).
		Where("title = ?title").
		OnConflict("DO NOTHING"). // optional
		SelectOrInsert()

	util.HandError("insert", err)
	fmt.Println(b)
	// 1. SELECT * FROM "books" WHERE title = 'my title'
	// 2. INSERT INTO "books" (title, text) VALUES ('my title', 'my text') RETURNING "id"
	// 3. go to step 1 on error
}

func insertMultiple() {
	b1 := &Book{AuthorId: "12345", Id: 4, Text: "insert test", Title: "insert"}
	b2 := &Book{AuthorId: "12345", Id: 5, Text: "insert test2", Title: "insert2"}

	result, err := db.Model(b1, b2).Insert()
	util.HandError("insert", err)
	fmt.Println(result)
	// INSERT INTO "books" (title, text) VALUES ('title1', 'text2'), ('title2', 'text2') RETURNING *
}

func insertReturningAll() {
	b := &Book{AuthorId: "12345", Id: 5, Text: "insert test2", Title: "insert2"}
	result, err := db.Model(b).Returning("*").Insert()
	util.HandError("insert", err)
	fmt.Println(result)
	// INSERT INTO "books" (title, text) VALUES ('my title', 'my text') RETURNING *
}

func insertReturningDefault() {
	b := &Book{AuthorId: "12345", Id: 4, Text: "insert test", Title: "insert"}
	// 默认返回ID
	err := db.Insert(b)
	util.HandError("insert", err)
	fmt.Println(b)
	// INSERT INTO "books" (title, text) VALUES ('my title', 'my text') RETURNING "id"
}

func ExampleColumn() {
	// Relation 3种用法
	// Relation("Author") 查出所有信息
	// Relation("Author.id") 查id
	// Relation("Author._") 只关联不查

	// Select book and associated author:
	// 查询book信息和作者信息
	columnRelation()

	// Select book id and associated author id:
	// 查询book信息和作者id

	columnRelation2()
	// Select book and join author without selecting it:
	// 只查book信息 不查询作者
	columnRelation3()
	// Join and select book author without selecting book:
	// 只查询作者信息 不查book
	columnRelation4()
}

func columnRelation4() {
	b := new(Book)
	err := db.Model(b).Column("_").Relation("Author").Select()
	util.HandError("select", err)
	fmt.Println("Relation", b)

	// SELECT "author"."id" AS "author__id", "author"."name" AS "author__name"
	// FROM "books"
	// LEFT JOIN "users" AS "author" ON "author"."id" = "book"."author_id"
	// WHERE id = 1
}

func columnRelation3() {
	b := new(Book)
	err := db.Model(b).Relation("Author._").Select()
	util.HandError("select", err)
	fmt.Println("Relation", b)

	// SELECT "book"."id"
	// FROM "books"
	// LEFT JOIN "users" AS "author" ON "author"."id" = "book"."author_id"
	// WHERE id = 1
}

func columnRelation2() {
	b := new(Book)
	err := db.Model(b).Column("book.id").Relation("Author.id").Select()
	util.HandError("select", err)
	fmt.Println("Relation", b)

	// SELECT "book"."id", "author"."id" AS "author__id"
	// FROM "books"
	// LEFT JOIN "users" AS "author" ON "author"."id" = "book"."author_id"
	// WHERE id = 1
}

func columnRelation() {
	b := new(Book)
	err := db.Model(b).Relation("Author").Select()
	util.HandError("select", err)
	fmt.Println("Relation", b)

	// SELECT
	//   "book"."id", "book"."title", "book"."text",
	//   "author"."id" AS "author__id", "author"."name" AS "author__name"
	// FROM "books"
	// LEFT JOIN "users" AS "author" ON "author"."id" = "book"."author_id"
	// WHERE id = 1
}

func ExampleCTE() {
	// CTE and subqueries
	// Select books using WITH statement:
	mCTEwith()
	// ame query using WrapWith:
	mCTEwrapWith()
	// Same query using subquery in select:
	mCTEsubQueries()
}

func mCTEsubQueries() {
	b := new(Book)
	authorBooks := db.Model((*Book)(nil)).Where("author_id = ?", 1)
	err := db.Model(nil).TableExpr("(?)", authorBooks).Select(b)
	util.HandError("wrapWith", err)
	fmt.Println(b)
	// SELECT * FROM (
	//   SELECT "book"."id", "book"."title", "book"."text"
	//   FROM "books" AS "book" WHERE (author_id = 1)
	// )
}

func mCTEwrapWith() {
	b := new(Book)
	err := db.Model(b).
		Where("author_id = ?", 1).
		WrapWith("author_books").
		Table("author_books").
		Select(b)
	util.HandError("wrapWith", err)
	fmt.Println(b)
	// WITH "author_books" AS (
	//   SELECT "book"."id", "book"."title", "book"."text"
	//   FROM "books" AS "book" WHERE (author_id = 1)
	// )
	// SELECT * FROM "author_books"
}

func mCTEwith() {
	b := new(Book)
	authorBooks := db.Model((*Book)(nil)).Where("author_id = ?", 1)
	err := db.Model().
		With("author_books", authorBooks).
		Table("author_books").
		Select(&b)
	util.HandError(" CTE with", err)
	fmt.Println(b)
	// WITH "author_books" AS (
	//   SELECT "book"."id", "book"."title", "book"."text"
	//   FROM "books" AS "book" WHERE (author_id = 1)
	// )
	// SELECT * FROM "author_books"
}

func ExampleJoin() {
	// Select book and manually join author:
	joinManually()
}

func joinManually() {
	b := new(Book)
	err := db.Model(b).
		ColumnExpr("book.*").
		ColumnExpr("a.id AS author__id, a.name AS author__name").
		Join("JOIN authors AS a ON a.id = book.author_id").
		First()
	util.HandError("select", err)
	fmt.Println("select", b)
	// SELECT book.*, a.id AS author__id, a.name AS author__name
	// FROM books
	// JOIN authors AS a ON a.id = book.author_id
	// ORDER BY id LIMIT 1

	//  join条件可以提取出来单独写
	// 	q.Join("LEFT JOIN authors AS a").
	// 	JoinOn("a.id = book.author_id").
	// 	JoinOn("a.active = ?", true)
}

func ExampleSelect() {
	fmt.Println("---------------ExampleSelect-----------------")
	//  根据主键查询
	selectByPkey()
	// 明确指定条件
	selectByParam()
	// 指定查询字段
	selectField()
	// 将查询结果存入变量
	select2Var()
	// Select book using WHERE ... OR ...:
	selectWhereOr()
	// Select book user WHERE ... AND (... OR ...):
	selectWhereAndOr()
	// Select first 20 books:
	selectOrderLimit()
	// Count books:
	selectCount()
	// Select 20 books and count all books:
	selectLimitCountAll()

	// Select by multiple ids:
	selectIn()
	// Select books for update
	selectForUpdate()
}

func selectForUpdate() {
	book := new(Book)
	ids := []int{1, 2, 3}
	err := db.Model(book).
		Where("id = ?", 1).
		For("UPDATE").
		Select()
	util.HandError("select", err)
	fmt.Println("select", ids)
	// SELECT * FROM books WHERE id  = 1 FOR UPDATE
}

func selectIn() {
	ids := []int{1, 2, 3}
	err := db.Model((*Book)(nil)).
		Where("id in (?)", pg.In(ids)).
		Select()
	util.HandError("select", err)
	fmt.Println("select", ids)
	// SELECT * FROM books WHERE id IN (1, 2, 3)
}

// 查询前20个统计所有个数
func selectLimitCountAll() {
	b := new(Book)
	count, err := db.Model(b).Limit(20).SelectAndCount()
	util.HandError("select", err)
	fmt.Println("count", count)
	fmt.Println("select", b)
	// SELECT "book"."id", "book"."title", "book"."text"
	// FROM "books" LIMIT 20
	//
	// SELECT count(*) FROM "books"
}

func selectCount() {
	count, err := db.Model((*Book)(nil)).Count()
	util.HandError("select", err)
	fmt.Println("count", count)
	// SELECT count(*) FROM "books"
}

func selectOrderLimit() {
	var books []Book
	err := db.Model(&books).Order("id ASC").Limit(20).Select()
	util.HandError("select", err)
	fmt.Println("select", books)
	// SELECT "book"."id", "book"."title", "book"."text"
	// FROM "books"
	// ORDER BY id ASC LIMIT 20
}

func selectWhereAndOr() {
	b := new(Book)
	err := db.Model(b).
		Where("title LIKE ?", "book%").
		WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			q = q.WhereOr("id = 1").
				WhereOr("id = 2")
			return q, nil
		}).
		Limit(1).
		Select()
	util.HandError("select", err)
	fmt.Println("select", b)
	// SELECT "book"."id", "book"."title", "book"."text"
	// FROM "books"
	// WHERE (title LIKE 'my%') AND (id = 1 OR id = 2)
	// LIMIT 1
}

func selectWhereOr() {
	b := new(Book)
	err := db.Model(b).
		Where("id = ?", "1").
		WhereOr("title LIKE ?", "f%").
		Limit(1).
		Select()
	util.HandError("select", err)
	fmt.Println("select", b)
	// SELECT "book"."id", "book"."title", "book"."text"
	// FROM "books"
	// WHERE (id > 100) OR (title LIKE 'my%')
	// LIMIT 1
}

func select2Var() {
	var title, text string
	err := db.Model((*Book)(nil)).
		Column("title", "text").
		Where("id = ?", 1).
		Select(&title, &text)
	util.HandError("select", err)
	fmt.Println("select", title, text)
	// SELECT "title", "text"
	// FROM "books" WHERE id = 1
}

func selectField() {
	b := new(Book)
	err := db.Model(b).Column("title", "text").Where("id = ?", 1).Select()
	util.HandError("select", err)
	fmt.Println("select", *b)
	// SELECT "title", "text"
	// FROM "books" WHERE id = 1
}

func selectByParam() {
	b := new(Book)
	err := db.Model(b).Where("id = ?", 1).Select()
	util.HandError("select", err)
	fmt.Println("select", *b)
	// SELECT "book"."id", "book"."title", "book"."text"
	// FROM "books" WHERE id = 1
}

//  根据主键查询
func selectByPkey() {
	b := &Book{Id: 1}
	err := db.Select(b)
	util.HandError("select", err)
	fmt.Println("select", *b)
	// SELECT "book"."id", "book"."title", "book"."text"
	// FROM "books" WHERE id = 1
}

func ExampleUpdate() {
	// Update all columns except primary keys:

	updateExceptPkey()

	// Update only column "title":
	updateOnly()
	updateOnly2()
	// Upper column "title" and scan it:

	updateUpper()

	// Update multiple books with single query:

	updateMultiple()
}

func updateMultiple() {
	b1 := &Book{AuthorId: "12345", Id: 5, Text: "insert test2", Title: "insert2"}
	b2 := &Book{AuthorId: "123455", Id: 6, Text: "insert test22", Title: "insert22"}
	result, err := db.Model(b1, b2).Update()
	util.HandError("updateMultiple ", err)
	fmt.Println(result)
	// UPDATE books AS book SET title = _data.title, text = _data.text
	// FROM (VALUES (1, 'title1', 'text1'), (2, 'title2', 'text2')) AS _data (id, title, text)
	// WHERE book.id = _data.id
}

// title替换为大写并返回
func updateUpper() {
	var title string
	b := &Book{AuthorId: "12345", Id: 5, Text: "insert test2", Title: "insert2"}
	_, err := db.Model(b).
		Set("title = upper(title)").
		Where("id = ?", 1).
		Returning("title").
		Update(&title)
	util.HandError("update", err)
	fmt.Println(title)

	// UPDATE books SET title = upper(title) WHERE id = 1 RETURNING title
}

func updateOnly2() {
	b := &Book{AuthorId: "12345", Id: 5, Text: "insert test2", Title: "insert2"}
	//  WherePK()=Where("id = ?id")
	res, err := db.Model(b).Column("title").WherePK().Update()
	// UPDATE books SET title = 'my title' WHERE id = 1
	util.HandError("update", err)
	fmt.Println(res)
	// UPDATE books SET title = 'my title' WHERE id = 1
}
func updateOnly() {
	b := &Book{AuthorId: "12345", Id: 5, Text: "insert test2", Title: "insert2"}
	res, err := db.Model(b).Set("title = ?title").Where("id = ?id").Update()
	util.HandError("update", err)
	fmt.Println(res)
	// UPDATE books SET title = 'my title' WHERE id = 1
}

func updateExceptPkey() {
	b := &Book{AuthorId: "12345", Id: 5, Text: "insert test2", Title: "insert2"}
	err := db.Update(b)
	util.HandError("update", err)
	// UPDATE books SET title = 'my title', text = 'my text' WHERE id = 1
}

func ExampleQuery() {
	fmt.Println("---------------ExampleQuery-----------------")

	var num int
	var err error
	// Simple params. 单个参数
	// 将根据query查询出的结果存入number
	_, err = db.Query(pg.Scan(&num), "SELECT id FROM Books WHERE name = ?", "first")
	if err != nil {
		panic(err)
	}
	fmt.Println("simple:", num)
	// Indexed params.根据位置取
	var name string
	_, err = db.Query(pg.Scan(&name), "SELECT name FROM Books WHERE id = ?0 AND age =?1", "1", 23)
	if err != nil {
		panic(err)
	}
	fmt.Println("indexed:", name)
	// Named params. 根据名字取
	params := &Params{
		X: 1,
		Y: 1,
	}
	_, err = db.Query(pg.Scan(&num), "SELECT ?x + ?y + ?Sum", params)
	if err != nil {
		panic(err)
	}
	fmt.Println("named:", num)
	// Model params. 直接根据模型
	var tableName, tableAlias, columns string
	_, err = db.Model(&Params{}).Query(
		pg.Scan(&tableName, &tableAlias, &columns),
		"SELECT '?TableName', '?TableAlias', '?Columns'",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("table name:", tableName)
	fmt.Println("table alias:", tableAlias)
	fmt.Println("columns:", columns)
}

func initTable() {
	err := db.CreateTable((*Author)(nil), &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	})
	insertAuthor()

	err = db.CreateTable((*Book)(nil), &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	})
	insertBook()

	util.HandError("initTable", err)
}

func insertBook() {
	author := &Author{Id: 81, Name: "author1", BookId: 12}
	for i := 0; i < 10; i++ {
		b := &Book{Id: i + 1, Title: "first book", Text: "desc", AuthorId: "123", Author: author}
		_, err := db.Model(b).Insert()
		util.HandError("insert", err)
	}
	// INSERT INTO "books" ("id", "title") VALUES (100, 'my title')
	// ON CONFLICT (id) DO UPDATE SET title = 'title version #1'
}

func insertAuthor() {
	for i := 80; i < 80; i++ {
		author := &Author{Id: i + 1, Name: "author1", BookId: 5}
		_, err := db.Model(author).Insert()
		util.HandError("insert", err)
	}
	// INSERT INTO "books" ("id", "title") VALUES (100, 'my title')
	// ON CONFLICT (id) DO UPDATE SET title = 'title version #1'
}
