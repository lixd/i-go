package dto

import "i-go/demo/cmodel"

type OrderReq struct {
	ID     uint    `json:"id" form:"id"`
	UserID uint    `json:"userId" form:"userId"`
	Amount float64 `json:"amount" form:"amount"`
	cmodel.PageModel
}

type OrderResp struct {
	ID     uint    `json:"id"`
	UserID uint    `json:"userId"`
	Amount float64 `json:"amount"`
}
