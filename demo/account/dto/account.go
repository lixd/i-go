package dto

import "i-go/demo/cmodel"

type AccountInsertReq struct {
	Id     uint    `json:"id" form:"id"`
	UserId uint    `json:"userId" form:"userId"`
	Amount float64 `json:"amount" form:"amount"`
}

type AccountReq struct {
	AccountInsertReq
	cmodel.Page
}

type AccountList struct {
	Data []AccountResp `json:"data"`
	Page cmodel.Page   `json:"page"`
}

type AccountResp struct {
	Id     uint    `json:"id"`
	UserId uint    `json:"userId"`
	Amount float64 `json:"amount"`
}
