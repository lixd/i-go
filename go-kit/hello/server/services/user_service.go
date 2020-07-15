package services

type IUser interface {
	GetName(userId int) string
}

type UserServer struct {
}

func (us *UserServer) GetName(userId int) string {
	if userId == 999 {
		return "admin"
	}
	return "guest"
}
