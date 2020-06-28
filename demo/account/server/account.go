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
	FindList(req *cmodel.Page) *ret.Result
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
		Model:  gorm.Model{ID: req.Id},
		UserId: req.UserId,
		Amount: req.Amount,
	}
	err := a.Dao.Insert(&account)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "create account"}).Error(err)
		return ret.Fail("", "db error")
	}
	// response the item which created by request
	res := dto.AccountResp{
		Id:     req.Id,
		UserId: req.UserId,
		Amount: req.Amount,
	}
	return ret.Success(&res)
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
		Model:  gorm.Model{ID: req.Id},
		UserId: req.UserId,
		Amount: req.Amount,
	}
	err := a.Dao.Update(&account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ret.Fail(err.Error())
		}
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新账户"}).Error(err)
		return ret.Fail("")
	}
	res := dto.AccountResp{
		Id:     req.Id,
		UserId: req.UserId,
		Amount: req.Amount,
	}
	return ret.Success(&res)
}

func (a *account) FindByUserId(userId uint) *ret.Result {
	res, err := a.Dao.FindByUserId(userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller()}).Error(err)
		return ret.Fail("", "db error")
	}
	account := dto.AccountResp{
		Id:     res.ID,
		UserId: res.UserId,
		Amount: res.Amount,
	}
	return ret.Success(&account)
}

func (a *account) FindList(req *cmodel.Page) *ret.Result {
	var resp dto.AccountList

	res, err := a.Dao.FindList(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	accounts := make([]dto.AccountResp, 0, len(res))
	var account dto.AccountResp
	for _, v := range res {
		account = dto.AccountResp{
			Id:     v.ID,
			UserId: v.UserId,
			Amount: v.Amount,
		}
		accounts = append(accounts, account)
	}
	resp.Data = accounts
	resp.Page = *req
	return ret.Success(&resp)
}
