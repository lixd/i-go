package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	amodel "i-go/demo/account/model"
	"i-go/demo/cmodel"
	"i-go/demo/order/dto"
	"i-go/demo/order/model"
)

type IOrder interface {
	Insert(req *model.Order) error
	DeleteById(id uint) error
	UpdateById(req *model.Order) error
	FindById(id uint) (*model.Order, error)
	Find(page *cmodel.PageModel) ([]model.Order, error)
	FindByUserId(req *dto.OrderReq) ([]model.Order, error)
	FindOrderAndUser() error
}

type order struct {
	DB *gorm.DB
}

func NewOrder(db *gorm.DB) IOrder {
	return &order{DB: db}
}

// Insert
func (o *order) Insert(req *model.Order) error {
	return o.DB.Transaction(func(tx *gorm.DB) error {

		// 检查金额
		var account amodel.Account
		cmd := tx.Where("user_id = ?", req.UserId).Find(&account)
		if err := cmd.Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if account.Amount < req.Amount {
			return errors.New("账户金额不足")
		}

		// 增加订单
		if err := tx.Create(req).Error; err != nil {
			return err
		}

		// 扣除金额
		cmd = tx.Model(amodel.Account{}).Where("user_id = ?", req.UserId).
			Update("amount", gorm.Expr("amount - ?", req.Amount))
		if err := cmd.Error; err != nil {
			return err
		}

		return nil
	})
}

// Delete
func (o *order) DeleteById(id uint) error {
	return o.DB.Delete(model.Order{}, "id = ? ", id).Error
}

// UpdateById
func (o *order) UpdateById(user *model.Order) error {
	return o.DB.Model(&model.Order{}).Where("id = ? ", user.ID).Update(user).Error
	//return u.DB.Model(&model.User{}).Where("id = ? ", user.ID).Update("name",user.Name,"age",user.Age).Error
	//return u.DB.Model(&model.User{}).Where("id = ? ", user.ID).Update(map[string]interface{}{"name":user.Name,"age":user.Age}).Error
}

// FindById
func (o *order) FindById(id uint) (*model.Order, error) {
	var order model.Order
	err := o.DB.Where("id = ? ", id).Find(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &model.Order{}, nil
		}
		return &model.Order{}, err
	}
	return &order, nil
}

// Find
func (o *order) Find(page *cmodel.PageModel) ([]model.Order, error) {

	users := make([]model.Order, 0, page.Size)
	err := o.DB.Model(&model.Order{}).Offset((page.Page - 1) * page.Size).Limit(page.Size).Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users, nil
		}
		return users, err
	}
	return users, nil
}

// FindByUserId
func (o *order) FindByUserId(req *dto.OrderReq) ([]model.Order, error) {
	users := make([]model.Order, 0, req.Size)
	err := o.DB.Model(&model.Order{}).Where("userId = ?", req.UserID).Offset((req.Page - 1) * req.Size).
		Limit(req.Size).Find(&users).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users, nil
		}
		return users, err
	}
	return users, nil
}

// FindOrderAndUser
func (o *order) FindOrderAndUser() error {
	var out [][]interface{}
	o.DB.Table("x_orders").Select("*").
		Joins("JOIN x_users as u ON u.id = x_orders.user_id").Find(&out)

	fmt.Println(out)
	//if err != nil {
	//	fmt.Println("error: ", err)
	//	if err == gorm.ErrRecordNotFound {
	//		return nil
	//	}
	//	return err
	//}
	//for cursor.Next() {
	//	var out []interface{}
	//	err:=cursor(&out)
	//	if err!=nil {
	//		fmt.Println("Scan error: ", err)
	//	}
	//	fmt.Println(out)
	//}
	return nil
}

func (o *order) txn() error {
	return o.DB.Transaction(func(tx *gorm.DB) error {
		return nil
	})
}
