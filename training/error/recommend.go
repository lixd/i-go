package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// 推荐的错误处理方式
// 参考 kratos https://github.com/go-kratos/kratos/blob/main/errors/errors.go
type Girl struct {
}

// 注意点1：这个错误定义在其他包，统一定义,不应该定义在 dao 层，这里demo就简单实现
var ErrNotFound = errors.New("no rows")

// dao 层的一个查询方法
func BatchGetGirl() ([]Girl, error) {
	db := sql.DB{}
	sqlStr := "SELECT * FROM grils WHERE love = 10"
	rows, err := db.Query(sqlStr)
	if err != nil {
		// 	handle error
	}
	err = rows.Err()
	if err != nil {
		// opqure sql. ErrNoRows
		// busssines code
		// stack stace
		// 注意点:2:这里将 dao 层的错误吞掉，返回外部统一定义的业务错误码
		return nil, errors.Wrapf(ErrNotFound, fmt.Sprintf("query: %s failed(%v)", sqlStr, err))
	}
	return []Girl{}, nil
}

// biz
func Usecase() error {
	v, err := BatchGetGirl()
	// Is errors 1.13 Unwrap, => root cause error,
	// bussiness code
	// sql or mongodb or hbase
	// 注意点3:biz层调用dao后处理错误时直接调用 is 方法进行判断即可，同时错误码由外部统一定义而不是定义在dao 层，所以 biz 层也不需要依赖 dao 层
	if errors.Is(ErrNotFound, err) {
	}
	fmt.Println(v)
	return nil
}
