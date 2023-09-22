package mock

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueryUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDB := NewMockIUser(mockCtrl)

	t.Run("empty id", func(t *testing.T) {
		mockDB.EXPECT().Get("").Return(User{}, ErrEmptyID).Times(0)
		user, err := QueryUser(mockDB, "")
		require.Equal(t, err, ErrEmptyID)
		require.Empty(t, user)
	})

	t.Run("normal id", func(t *testing.T) {
		targetUser := User{
			Username: "tom",
			Password: "pwd",
		}
		mockDB.EXPECT().Get("tom").Return(targetUser, nil).Times(1)
		user, err := QueryUser(mockDB, "tom")
		require.NoError(t, err)
		require.Equal(t, user, user)
	})
	t.Run("not exist user", func(t *testing.T) {
		mockDB.EXPECT().Get("not_exist").Return(User{}, ErrUserNotFond).Times(1)
		user, err := QueryUser(mockDB, "not_exist")
		require.Equal(t, err, ErrUserNotFond)
		require.Empty(t, user)
	})
}
