package model

// Image
type Image struct {
	Code  string `json:"code"`
	Msg   string `json:"msg"`
	ImgId string `json:"imgid"`
	Knock string `json:"knock"`
}

// Validate
type Validate struct {
	Code  string `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

// Verify
type Verify struct {
	Msg     string `json:"msg"`     // 错误信息
	Success int    `json:"success"` // 验证结果，1为通过，0为失败
	Score   int    `json:"score"`   // 可信度，区间[0, 100] 离线模式默认为100
}
