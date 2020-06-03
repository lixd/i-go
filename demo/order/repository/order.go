package repository

import (
	"github.com/jinzhu/gorm"
	"i-go/demo/cmodel"
	"i-go/demo/order/model"
)

type IOrder interface {
	Insert(req *model.Order) error
	DeleteById(id uint) error
	UpdateById(req *model.Order) error
	FindById(id uint) (*model.Order, error)
	Find(page *cmodel.PageModel) ([]model.Order, error)
}

type order struct {
	DB *gorm.DB
}

func NewOrder(db *gorm.DB) IOrder {
	return &order{DB: db}
}

// Insert
func (o *order) Insert(req *model.Order) error {
	return o.DB.Create(req).Error
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
