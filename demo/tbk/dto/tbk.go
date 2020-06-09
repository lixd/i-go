package dto

type TBKReq struct {
	KeyWords string `form:"keyWords"`
}

type TBKItem struct {
	ShareURL string `json:"shareURL"`
}
type TBKResp struct {
	List []TBKItem `json:"list"`
}
