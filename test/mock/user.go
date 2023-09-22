package mock

import (
	"github.com/pkg/errors"
)

//go:generate mockgen -source=user.go -destination=mock_user.go -package mock i-go/test/mock IUser
type IUser interface {
	Get(id string) (User, error)
}
type User struct {
	Username string
	Password string
}

var (
	ErrEmptyID     = errors.New("id is empty")
	ErrUserNotFond = errors.New("user not found")
)

func QueryUser(db IUser, id string) (User, error) {
	if id == "" {
		return User{}, ErrEmptyID
	}
	return db.Get(id)
}
