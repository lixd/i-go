package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	amodel "i-go/demo/account/model"
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

// Insert 自动处理事务
func (u *user) Insert(req *model.User) error {
	return u.DB.Transaction(func(tx *gorm.DB) error {
		// 1.创建 用户
		cmd := tx.Create(req) // 这里面必须用tx 而不是DB 否则就和没事务一样...
		if err := cmd.Error; err != nil {
			return err
		}
		user, ok := cmd.Value.(*model.User)
		if !ok {
			fmt.Println("断言失败: ", cmd.Value)
			return errors.New("txn error")
		}
		fmt.Printf("%#v \n", user)

		// 2. 根据 userId 创建账户
		var account = amodel.Account{
			UserId: user.ID,
			Amount: 0.0,
		}
		cmd = tx.Create(&account)
		if err := cmd.Error; err != nil {
			return err
		}
		return nil
	})
}

// InsertCustom 手动处理事务
func (u *user) InsertCustom(req *model.User) error {
	tx := u.DB.Begin()
	var err error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		return err
	}
	// 1.创建 用户
	cmd := tx.Create(req)

	if err = cmd.Error; err != nil {
		tx.Rollback()
		return err
	}

	user, ok := cmd.Value.(*model.User)
	if !ok {
		fmt.Println("断言失败: ", cmd.Value)
		tx.Rollback()
		return errors.New("txn error")
	}
	fmt.Printf("%#v \n", user)

	// 2. 根据 userId 创建账户
	var account = amodel.Account{
		UserId: user.ID,
		Amount: 0.0,
	}
	cmd = tx.Create(&account)
	if err = cmd.Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// DeleteById
func (u *user) DeleteById(userId uint) error {
	return u.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 检查账户金额
		var account amodel.Account
		cmd := tx.Model(amodel.Account{}).Where("user_id = ? ", userId).Find(&account)
		if err := cmd.Error; err != nil {
			return err
		}
		if account.Amount != 0.0 {
			return errors.New("请先清理账户余额")
		}
		// 2. 删除账户
		cmd = tx.Where("user_id = ?", account.UserId).Delete(amodel.Account{})
		if err := cmd.Error; err != nil {
			return err
		}
		// 3. 删除用户
		cmd = tx.Where("id = ?", userId).Delete(model.User{})
		if err := cmd.Error; err != nil {
			return err
		}
		return nil
	})
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
