package server

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"i-go/demo/account/dto"
	"i-go/demo/account/model"
	"i-go/demo/account/repository"
	"i-go/demo/cmodel"
	"i-go/demo/ret"

	"i-go/utils"
)

type IAccount interface {
	Insert(req *dto.AccountInsertReq) *ret.Result
	DeleteByUserId(userId uint) *ret.Result
	Update(req *dto.AccountReq) *ret.Result
	FindByUserId(userId uint) *ret.Result
	FindList(req *cmodel.PageModel) *ret.Result
}

type account struct {
	Dao repository.IAccount
}

func NewAccount(dao repository.IAccount) IAccount {
	return &account{Dao: dao}
}

// Insert
func (a *account) Insert(req *dto.AccountInsertReq) *ret.Result {
	account := model.Account{
		Model:  gorm.Model{ID: req.ID},
		UserId: uint(1),
		Amount: 12.11,
	}
	err := a.Dao.Insert(&account)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "create account"}).Error(err)
		return ret.Fail("", "db error")
	}
	// response the item which created by request
	return ret.Success(&account)
}

func (a *account) DeleteByUserId(userId uint) *ret.Result {
	err := a.Dao.DeleteByUserId(userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "删除用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	return ret.Success("")
}

func (a *account) Update(req *dto.AccountReq) *ret.Result {
	account := model.Account{
		Model:  gorm.Model{ID: req.ID},
		UserId: req.UserID,
		Amount: req.Amount,
	}
	err := a.Dao.Update(&account)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	return ret.Success(&account)
}

func (a *account) FindByUserId(userId uint) *ret.Result {
	res, err := a.Dao.FindByUserId(userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新账户"}).Error(err)
		return ret.Fail("", "db error")
	}
	account := dto.AccountResp{
		ID:     res.ID,
		UserID: res.UserId,
		Amount: res.Amount,
	}
	return ret.Success(&account)
}

func (a *account) FindList(req *cmodel.PageModel) *ret.Result {
	res, err := a.Dao.FindList(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	accounts := make([]dto.AccountResp, 0, len(res))
	var account dto.AccountResp
	for _, v := range res {
		account = dto.AccountResp{
			ID:     v.ID,
			UserID: v.UserId,
			Amount: v.Amount,
		}
		accounts = append(accounts, account)
	}
	return ret.Success(&accounts)
}
