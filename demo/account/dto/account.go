package dto

import "i-go/demo/cmodel"

type AccountInsertReq struct {
	ID     uint    `json:"id" form:"id"`
	UserID uint    `json:"userId" form:"userId"`
	Amount float64 `json:"amount" form:"amount"`
}

type AccountReq struct {
	AccountInsertReq
	cmodel.PageModel
}

type AccountResp struct {
	ID     uint    `json:"id"`
	UserID uint    `json:"userId"`
	Amount float64 `json:"amount"`
}
