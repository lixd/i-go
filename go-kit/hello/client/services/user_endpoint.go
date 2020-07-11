package services

type UserRequest struct {
	UserId int `json:"userId"`
}
type UserResponse struct {
	Username string `json:"username"`
}
