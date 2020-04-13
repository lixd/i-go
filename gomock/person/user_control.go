package person

type userControl struct {
	IUC IUserControl
}

func NewUserControl(p IUserControl) *userControl {
	return &userControl{IUC: p}
}

func (uc *userControl) Login(username, password string) error {
	return uc.IUC.Login(username, password)
}
