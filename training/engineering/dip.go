package main

import (
	"database/sql"
)

// 依赖倒置原则（Dependence Inversion Principle）
// 依赖注入

// IService 上层只依赖接口(抽象)而不依赖实现(具象)
type IService interface {
	Query(id int) (string, error)
}

type Service struct {
	db *sql.DB
}

// NewService 通道外部传入db对象来实现依赖倒置
func NewService(db *sql.DB) IService {
	return &Service{db: db}
}

func (s *Service) Query(id int) (string, error) {
	// s.db.Query()
	return "", nil
}
