package server

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"i-go/core/logger/izap"
	"i-go/demo/cmodel"
	"i-go/demo/order/dto"
	"i-go/demo/order/model"
	"i-go/demo/order/repository"
	"i-go/demo/ret"
	"i-go/utils"
)

type IOrder interface {
	Insert(req *dto.OrderReq) *ret.Result
	DeleteById(req *dto.OrderReq) *ret.Result
	UpdateById(req *dto.OrderReq) *ret.Result
	FindById(req *dto.OrderReq) *ret.Result
	Find(req *cmodel.PageModel) *ret.Result
	FindOrderAndUser() *ret.Result
}

type order struct {
	Dao repository.IOrder
}

func NewOrder(dao repository.IOrder) IOrder {
	return &order{Dao: dao}
}

func (o *order) Insert(req *dto.OrderReq) *ret.Result {
	user := model.Order{
		Model:  gorm.Model{ID: req.ID},
		UserId: uint(1),
		Amount: 12.11,
	}
	err := o.Dao.Insert(&user)
	if err != nil {
		//logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "新增订单"}).Error(err)
		//izap.Logger.Info(zap.String("scenes", "新增订单"), zap.String("error", err.Error()))
		izap.Logger.Infof("scenes:%s,error:%v", "新增订单", err.Error())
		return ret.Fail("", err.Error())
	}
	return ret.Success("")
}

func (o *order) DeleteById(req *dto.OrderReq) *ret.Result {
	err := o.Dao.DeleteById(req.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "删除用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	return ret.Success("")
}

func (o *order) UpdateById(req *dto.OrderReq) *ret.Result {
	user := model.Order{
		Model:  gorm.Model{ID: req.ID},
		UserId: req.UserID,
		Amount: req.Amount,
	}
	err := o.Dao.UpdateById(&user)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	return ret.Success("")
}

func (o *order) FindById(req *dto.OrderReq) *ret.Result {
	res, err := o.Dao.FindById(req.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	order := dto.OrderResp{
		ID:     res.ID,
		UserID: res.UserId,
		Amount: res.Amount,
	}
	return ret.Success(&order)
}

func (o *order) Find(req *cmodel.PageModel) *ret.Result {
	res, err := o.Dao.Find(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	users := make([]dto.OrderResp, 0, len(res))
	var user dto.OrderResp
	for _, v := range res {
		user = dto.OrderResp{
			ID:     v.ID,
			UserID: v.UserId,
			Amount: v.Amount,
		}
		users = append(users, user)
	}
	return ret.Success(&users)
}
func (o *order) FindByUserId(req *dto.OrderReq) *ret.Result {
	res, err := o.Dao.FindByUserId(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	users := make([]dto.OrderResp, 0, len(res))
	var user dto.OrderResp
	for _, v := range res {
		user = dto.OrderResp{
			ID:     v.ID,
			UserID: v.UserId,
			Amount: v.Amount,
		}
		users = append(users, user)
	}
	return ret.Success(&users)
}
func (o *order) FindOrderAndUser() *ret.Result {
	err := o.Dao.FindOrderAndUser()
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return ret.Fail("", "db error")
	}
	return ret.Success("")
}
