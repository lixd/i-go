package person

// 批量生成
//go:generate mockgen -destination=./user_control_mock.go -package=person . IUserControl

type IUserControl interface {
	Login(username, password string) error
}
