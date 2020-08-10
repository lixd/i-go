package model

type Login struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Token    string `form:"token"`
}
type Offline struct {
	Action   string `form:"offline_action"`
	Callback string `form:"callback"`
	Knock    string `form:"knock"`
	UserCode string `form:"v"`
}
