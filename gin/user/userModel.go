package user

type User struct {
	Name    string `json:"name" form:"name"`
	Age     int    `json:"age" form:"age"`
	Address string `json:"address" form:"address"`
}
type LoginForm struct {
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
}
type Class struct {
	Id     string `json:"id" form:"id"`
	Number int    `json:"number" form:"number"`
}
