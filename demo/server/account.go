package server

import (
	"errors"
	resultcode2 "i-go/demo/constant/resultcode"
	"i-go/demo/model"
	"i-go/demo/repository"
	"i-go/demo/utils"
)

type Server interface {
	LoginServer(phone, password string) (result *model.Response, err error)
	RegisterServer(phone, password string) (result *model.Response, err error)
}
type AccountServer struct {
	Dao *repository.DAO
}

func (as *AccountServer) LoginServer(phone, password string) (result *model.Response, err error) {
	phoneVerify := utils.PhoneVerify(phone)
	if !phoneVerify {
		return &model.Response{
			Code: resultcode2.Fail,
			Msg:  "phone or password verify error"}, errors.New("phone or password verify error")
	}
	// user, err := repository.ADAO.FindUserByPhone(phone)
	user, err := (*as.Dao).FindUserByPhone(phone)
	if err != nil {
		return &model.Response{Code: resultcode2.Fail, Msg: "用户不存在。"}, err
	}
	return &model.Response{
		Code: resultcode2.Success,
		Data: &user,
		Msg:  "success"}, nil
}
func (as *AccountServer) RegisterServer(phone, password string) (result *model.Response, err error) {
	phoneVerify := utils.PhoneVerify(phone)
	if !phoneVerify {
		return &model.Response{
			Code: resultcode2.Fail,
			Msg:  "phone or password verify error"}, errors.New("phone or password verify error")
	}
	user, err := (*as.Dao).FindUserByPhone(phone)
	if err == nil {
		return &model.Response{Code: resultcode2.Fail, Msg: err.Error()}, err
	}
	if phone == user.Phone {
		return &model.Response{Code: resultcode2.Fail, Msg: "用户已存在。"}, err
	}
	userId, err := (*as.Dao).CreateUser(phone, password)
	if err != nil {
		return &model.Response{
			Code: resultcode2.Fail,
			Msg:  err.Error()}, nil
	}
	return &model.Response{
		Code: resultcode2.Fail,
		Data: userId,
		Msg:  "success"}, nil
}
