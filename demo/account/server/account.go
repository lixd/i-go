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
	Insert(req *dto.AccountReq) *ret.Result
	DeleteByUserId(req *dto.AccountReq) *ret.Result
	UpdateById(req *dto.AccountReq) *ret.Result
	FindByUserId(req *dto.AccountReq) *ret.Result
	Find(req *cmodel.PageModel) *ret.Result
}

type account struct {
	Dao repository.IAccount
}

func NewAccount(dao repository.IAccount) IAccount {
	return &account{Dao: dao}
}

func (a *account) Insert(req *dto.AccountReq) *ret.Result {
	user := model.Account{
		Model:  gorm.Model{ID: req.ID},
		UserId: uint(1),
		Amount: 12.11,
	}
	err := a.Dao.Insert(&user)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "新增用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	return ret.Success("")
}

func (a *account) DeleteByUserId(req *dto.AccountReq) *ret.Result {
	err := a.Dao.DeleteByUserId(req.UserID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "删除用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	return ret.Success("")
}

func (a *account) UpdateById(req *dto.AccountReq) *ret.Result {
	user := model.Account{
		Model:  gorm.Model{ID: req.ID},
		UserId: req.UserID,
		Amount: req.Amount,
	}
	err := a.Dao.UpdateByUserId(&user)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	return ret.Success("")
}

func (a *account) Find(req *cmodel.PageModel) *ret.Result {
	res, err := a.Dao.Find(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	users := make([]dto.AccountResp, 0, len(res))
	var user dto.AccountResp
	for _, v := range res {
		user = dto.AccountResp{
			ID:     v.ID,
			UserID: v.UserId,
			Amount: v.Amount,
		}
		users = append(users, user)
	}
	return ret.Success(&users)
}

func (a *account) FindByUserId(req *dto.AccountReq) *ret.Result {
	res, err := a.Dao.FindByUserId(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	account := dto.AccountResp{
		ID:     res.ID,
		UserID: res.UserId,
		Amount: res.Amount,
	}
	return ret.Success(&account)
}
