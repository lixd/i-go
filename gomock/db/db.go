package db

type DB interface {
	Get(key string) (int, error)
}

// mockgen 生成 mock 文件 一般传递三个参数。包含需要被mock的接口得到源文件source，生成的目标文件destination，包名package

//go:generate mockgen -source=db.go -destination=./db_mock.go -package=db
func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}
