package repository

import (
	"github.com/jinzhu/gorm"
	"i-go/demo/cmodel"
	"i-go/demo/user/model"
)

type IUser interface {
	Insert(req *model.User) error
	DeleteById(id uint) error
	UpdateById(req *model.User) error
	FindById(id uint) (*model.User, error)
	Find(page *cmodel.PageModel) ([]model.User, error)
}

type user struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) IUser {
	return &user{DB: db}
}

// Insert
func (u *user) Insert(req *model.User) error {
	return u.DB.Create(req).Error
}

// Delete
func (u *user) DeleteById(id uint) error {
	return u.DB.Delete(model.User{}, "id = ? ", id).Error
}

// UpdateById
func (u *user) UpdateById(user *model.User) error {
	return u.DB.Model(&model.User{}).Where("id = ? ", user.ID).Update(user).Error
	//return u.DB.Model(&model.User{}).Where("id = ? ", user.ID).Update("name",user.Name,"age",user.Age).Error
	//return u.DB.Model(&model.User{}).Where("id = ? ", user.ID).Update(map[string]interface{}{"name":user.Name,"age":user.Age}).Error
}

// FindById
func (u *user) FindById(id uint) (*model.User, error) {
	var user model.User
	err := u.DB.Where("id = ? ", id).Find(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &model.User{}, nil
		}
		return &model.User{}, err
	}
	return &user, nil
}

// Find
func (u *user) Find(page *cmodel.PageModel) ([]model.User, error) {

	users := make([]model.User, 0, page.Size)
	err := u.DB.Model(&model.User{}).Offset((page.Page - 1) * page.Size).Limit(page.Size).Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users, nil
		}
		return users, err
	}
	return users, nil
}
