package repository

import (
	"errors"
	"fmt"

	amodel "i-go/demo/account/model"
	"i-go/demo/cmodel"
	"i-go/demo/order/model"

	"gorm.io/gorm"
)

type IOrder interface {
	Insert(req *model.Order) error
	Delete(id uint) error
	Update(req *model.Order) error
	FindById(id uint) (*model.Order, error)
	Find(userId uint, page *cmodel.Page) ([]model.Order, error)
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

// DeleteById
func (o *order) Delete(id uint) error {
	return o.DB.Delete(model.Order{}, "id = ? ", id).Error
}

func (o *order) Update(order *model.Order) error {
	cmd := o.DB.Model(&model.Order{}).Where("id = ? ", order.ID).Updates(order)
	if err := cmd.Error; err != nil {
		return err
	}
	if cmd.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
	// return u.DB.Model(&model.User{}).Where("id = ? ", user.Id).Update("name",user.Name,"age",user.Age).Error
	// return u.DB.Model(&model.User{}).Where("id = ? ", user.Id).Update(map[string]interface{}{"name":user.Name,"age":user.Age}).Error
}

func (o *order) FindById(id uint) (*model.Order, error) {
	var order model.Order
	err := o.DB.Where("id = ? ", id).Find(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &order, nil
		}
		return &order, err
	}
	return &order, nil
}

func (o *order) Find(userId uint, page *cmodel.Page) ([]model.Order, error) {
	users := make([]model.Order, 0, page.Size)

	err := o.DB.Model(&model.Order{}).Count(&page.Total).Error
	if err != nil {
		return users, err
	}

	err = o.DB.Model(&model.Order{}).Offset(int(page.Skip())).Limit(page.Size).Find(&users).Error
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
	// if err != nil {
	//	fmt.Println("error: ", err)
	//	if err == gorm.ErrRecordNotFound {
	//		return nil
	//	}
	//	return err
	// }
	// for cursor.Next() {
	//	var out []interface{}
	//	err:=cursor(&out)
	//	if err!=nil {
	//		fmt.Println("Scan error: ", err)
	//	}
	//	fmt.Println(out)
	// }
	return nil
}
