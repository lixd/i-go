package repository

import (
	"errors"

	"i-go/demo/account/model"
	"i-go/demo/cmodel"
	umodel "i-go/demo/user/model"
	"i-go/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IAccount interface {
	Insert(req *model.Account) error
	DeleteByUserId(userId uint) error
	Update(req *model.Account) error
	FindByUserId(userId uint) (model.Account, error)
	FindList(page *cmodel.Page) ([]model.Account, error)
}

type account struct {
	DB *gorm.DB
}

func NewAccount(db *gorm.DB) IAccount {
	return &account{DB: db}
}

// Insert
func (a *account) Insert(req *model.Account) error {
	return a.DB.Transaction(func(tx *gorm.DB) error {
		// check user is exist
		var user umodel.User
		cmd := tx.Model(&user).Where("id = ?", req.UserId).Find(&user)
		if err := cmd.Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("invalid user")
			}
			return err
		}

		// create user
		if err := tx.Create(req).Error; err != nil {
			return err
		}

		return nil
	})
}

// DeleteByUserId
func (a *account) DeleteByUserId(userId uint) error {
	// return a.DB.Delete(model.Account{}, "user_id = ? ", userId).Error
	return a.DB.Where("user_id = ? ", userId).Delete(model.Account{}).Error
}

// Update
func (a *account) Update(account *model.Account) error {
	cmd := a.DB.Model(&model.Account{}).Where("user_id = ? ", account.UserId).
		Update("amount", account.Amount)

	if err := cmd.Error; err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新账户"}).Error(err)
		return err
	}
	if cmd.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// FindByUserId
func (a *account) FindByUserId(userId uint) (model.Account, error) {
	var account model.Account
	err := a.DB.Model(&account).Where("user_id = ?", userId).Find(&account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return account, nil
		}
		return account, err
	}
	return account, nil
}

// FindList
func (a *account) FindList(page *cmodel.Page) ([]model.Account, error) {
	users := make([]model.Account, 0, page.Size)
	// 查询总记录数并存储到 page 中
	err := a.DB.Model(&model.Account{}).Count(&page.Total).Error
	if err != nil {
		return users, err
	}
	err = a.DB.Model(&model.Account{}).Offset(int(page.Skip())).Limit(page.Size).Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users, nil
		}
		return users, err
	}
	return users, nil
}
