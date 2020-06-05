package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"i-go/demo/account/dto"
	"i-go/demo/account/model"
	"i-go/demo/cmodel"
	umodel "i-go/demo/user/model"
)

type IAccount interface {
	Insert(req *model.Account) error
	DeleteByUserId(id uint) error
	UpdateByUserId(req *model.Account) error
	Find(page *cmodel.PageModel) ([]model.Account, error)
	FindByUserId(req *dto.AccountReq) (model.Account, error)
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

		// 检查金额
		var user umodel.User
		cmd := tx.Where("user_id = ?", req.UserId).Find(&user)
		if err := cmd.Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("用户不存在")
			}
			return err
		}

		// 创建账户
		if err := tx.Create(req).Error; err != nil {
			return err
		}

		return nil
	})
}

// Delete
func (a *account) DeleteByUserId(userId uint) error {
	return a.DB.Delete(model.Account{}, "user_id = ? ", userId).Error
	//return a.DB.Where("user_id = ? ", userId).Delete(model.Account{}).Error
}

// UpdateById
func (a *account) UpdateByUserId(account *model.Account) error {
	return a.DB.Model(&model.Account{}).Where("user_id = ? ", account.UserId).
		Update("amount", account.Amount).Error
}

// Find
func (a *account) Find(page *cmodel.PageModel) ([]model.Account, error) {

	users := make([]model.Account, 0, page.Size)
	err := a.DB.Model(&model.Account{}).Offset((page.Page - 1) * page.Size).Limit(page.Size).Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users, nil
		}
		return users, err
	}
	return users, nil
}

// FindByUserId
func (a *account) FindByUserId(req *dto.AccountReq) (model.Account, error) {
	var account model.Account
	err := a.DB.Model(&model.Account{}).Where("user_id = ?", req.UserID).Find(&account).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return account, nil
		}
		return account, err
	}
	return account, nil
}
