package server

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"i-go/demo/cmodel"
	"i-go/demo/ret"
	"i-go/demo/user/dto"
	"i-go/demo/user/model"
	"i-go/demo/user/repository"
	"i-go/utils"
)

type IUser interface {
	Insert(req *dto.UserReq) *ret.Result
	DeleteById(req *dto.UserReq) *ret.Result
	UpdateById(req *dto.UserReq) *ret.Result
	FindById(req *dto.UserReq) *ret.Result
	Find(req *cmodel.PageModel) *ret.Result
}

type user struct {
	Dao repository.IUser
}

func NewUser(dao repository.IUser) IUser {
	return &user{Dao: dao}
}

func (u *user) Insert(req *dto.UserReq) *ret.Result {
	user := model.User{
		Model:      gorm.Model{ID: req.ID},
		Name:       req.Name,
		Pwd:        req.Pwd,
		Phone:      req.Phone,
		Age:        req.Age,
		RegisterIP: req.RegisterIP,
		LoginIP:    req.LoginIP,
	}
	err := u.Dao.Insert(&user)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "新增用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	return ret.Success("")
}

func (u *user) DeleteById(req *dto.UserReq) *ret.Result {
	err := u.Dao.DeleteById(req.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "删除用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	return ret.Success("")
}

func (u *user) UpdateById(req *dto.UserReq) *ret.Result {
	user := model.User{
		Model: gorm.Model{ID: req.ID},
		Name:  req.Name,
		Phone: req.Phone,
		Pwd:   req.Pwd,
		Age:   req.Age,
	}
	err := u.Dao.UpdateById(&user)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	return ret.Success("")
}

func (u *user) FindById(req *dto.UserReq) *ret.Result {
	res, err := u.Dao.FindById(req.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	user := dto.UserResp{
		ID:    res.ID,
		Name:  res.Name,
		Phone: req.Phone,
		Pwd:   res.Pwd,
		Age:   res.Age,
	}
	return ret.Success(&user)
}

func (u *user) Find(req *dto.PageModel) *ret.Result {
	res, err := u.Dao.Find(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	users := make([]dto.UserResp, 0, len(res))
	var user dto.UserResp
	for _, v := range res {
		user = dto.UserResp{
			ID:    v.ID,
			Name:  v.Name,
			Phone: v.Phone,
			Pwd:   v.Pwd,
			Age:   v.Age,
		}
		users = append(users, user)
	}
	return ret.Success(&users)
}
