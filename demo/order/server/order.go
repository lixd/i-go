package server

import (
	"i-go/demo/cmodel"
	"i-go/demo/common/ret/srv"
	"i-go/demo/order/dto"
	"i-go/demo/order/model"
	"i-go/demo/order/repository"
	"i-go/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IOrder interface {
	Insert(req *dto.OrderReq) *srv.Result
	Delete(req *dto.OrderReq) *srv.Result
	Update(req *dto.OrderReq) *srv.Result
	FindById(id uint) *srv.Result
	Find(req *dto.OrderReq) *srv.Result
	FindOrderAndUser() *srv.Result
}

type order struct {
	Dao repository.IOrder
}

func NewOrder(dao repository.IOrder) IOrder {
	return &order{Dao: dao}
}

func (o *order) Insert(req *dto.OrderReq) *srv.Result {
	order := model.Order{
		Model:  gorm.Model{ID: req.Id},
		UserId: req.UserId,
		Amount: req.Amount,
	}

	err := o.Dao.Insert(&order)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "新增订单"}).Error(err)
		return srv.Fail("", err.Error())
	}

	res := dto.OrderResp{
		Id:     req.Id,
		UserId: req.UserId,
		Amount: req.Amount,
	}
	return srv.Success(&res)
}

func (o *order) Delete(req *dto.OrderReq) *srv.Result {
	err := o.Dao.Delete(req.Id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "删除用户"}).Error(err)
		return srv.Fail("", "db error")
	}
	return srv.Success(&dto.OrderResp{})
}

func (o *order) Update(req *dto.OrderReq) *srv.Result {
	order := model.Order{
		Model:  gorm.Model{ID: req.Id},
		UserId: req.UserId,
		Amount: req.Amount,
	}
	err := o.Dao.Update(&order)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return srv.Fail(err.Error())
		}
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return srv.Fail("")
	}
	res := dto.OrderResp{
		Id:     req.Id,
		UserId: req.UserId,
		Amount: req.Amount,
	}
	return srv.Success(&res)
}

func (o *order) FindById(id uint) *srv.Result {
	res, err := o.Dao.FindById(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller()}).Error(err)
		return srv.Fail("", "db error")
	}
	order := dto.OrderResp{
		Id:     res.ID,
		UserId: res.UserId,
		Amount: res.Amount,
	}
	return srv.Success(&order)
}

func (o *order) Find(req *dto.OrderReq) *srv.Result {
	var resp dto.OrderList

	page := cmodel.NewPaging(req.Page.Page, req.Page.Size)
	res, err := o.Dao.Find(req.UserId, page)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return srv.Fail("", "db error")
	}
	users := make([]dto.OrderResp, 0, len(res))
	var user dto.OrderResp
	for _, v := range res {
		user = dto.OrderResp{
			Id:     v.ID,
			UserId: v.UserId,
			Amount: v.Amount,
		}
		users = append(users, user)
	}
	resp.Data = users
	resp.Page = *page
	return srv.Success(&resp)
}

func (o *order) FindOrderAndUser() *srv.Result {
	err := o.Dao.FindOrderAndUser()
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "更新用户"}).Error(err)
		return srv.Fail("", "db error")
	}
	return srv.Success("")
}
