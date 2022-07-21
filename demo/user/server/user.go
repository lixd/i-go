package server

import (
	"i-go/demo/cmodel"
	"i-go/demo/common/ret/srv"
	"i-go/demo/user/dto"
	"i-go/demo/user/model"
	"i-go/demo/user/repository"
	"i-go/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUser interface {
	Insert(req *dto.UserReq) *srv.Result
	DeleteById(req *dto.UserReq) *srv.Result
	UpdateById(req *dto.UserReq) *srv.Result
	FindById(id uint) *srv.Result
	Find(req *cmodel.Page) *srv.Result
}

type user struct {
	Dao repository.IUser
}

func NewUser(dao repository.IUser) IUser {
	return &user{Dao: dao}
}

func (u *user) Insert(req *dto.UserReq) *srv.Result {
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
	// err := u.Dao.InsertCustom(&user)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "新增用户"}).Error(err)
		return srv.Fail("", "db error")
	}
	return srv.Success("")
}

func (u *user) DeleteById(req *dto.UserReq) *srv.Result {
	err := u.Dao.DeleteById(req.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "删除用户"}).Error(err)
		return srv.Fail("", "db error")
	}
	return srv.Success("")
}

func (u *user) UpdateById(req *dto.UserReq) *srv.Result {
	user := model.User{
		Model: gorm.Model{ID: req.ID},
		Name:  req.Name,
		Phone: req.Phone,
		Pwd:   req.Pwd,
		Age:   req.Age,
	}
	err := u.Dao.UpdateById(&user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return srv.Fail(err.Error())
		}
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return srv.Fail("", "db error")
	}
	return srv.Success("")
}

func (u *user) FindById(id uint) *srv.Result {
	res, err := u.Dao.FindById(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return srv.Fail("", "db error")
	}
	user := dto.UserResp{
		ID:    res.ID,
		Name:  res.Name,
		Phone: res.Phone,
		Pwd:   res.Pwd,
		Age:   res.Age,
	}
	return srv.Success(&user)
}

func (u *user) Find(req *cmodel.Page) *srv.Result {
	var resp dto.UserList
	res, err := u.Dao.Find(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return srv.Fail("", "db error")
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
	resp.Data = users
	resp.Page = *req
	return srv.Success(&resp)
}
