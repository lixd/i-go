package dto

import "i-go/demo/cmodel"

type OrderReq struct {
	Id     uint    `json:"id" form:"id"`
	UserId uint    `json:"userId" form:"userId"`
	Amount float64 `json:"amount" form:"amount"`
	cmodel.Page
}

type OrderList struct {
	Data []OrderResp `json:"data"`
	Page cmodel.Page `json:"page"`
}
type OrderResp struct {
	Id     uint    `json:"id"`
	UserId uint    `json:"userId"`
	Amount float64 `json:"amount"`
}
