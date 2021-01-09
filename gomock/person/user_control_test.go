package person

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserControl_Login(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	var (
		username = "admin"
		password = "root"
	)

	mockUC := NewMockIUserControl(ctl)
	gomock.InOrder(
		mockUC.EXPECT().Login(username, password).Return(nil),
	)
	control := NewUserControl(mockUC)
	err := control.Login(username, password)
	if err != nil {
		t.Errorf("user.GetUserInfo err: %v", err)
	}
}
