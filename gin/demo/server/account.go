package server

import (
	"i-go/gin/demo/constant/resultcode"
	"i-go/gin/demo/model"
	"i-go/gin/demo/repository"
	"i-go/gin/demo/utils"
)

type Server interface {
	LoginServer(phone, password string) (result *model.Response, err error)
}
type AccountServer struct {
	DAO *repository.DAO
}

func (as *AccountServer) LoginServer(phone, password string) (result *model.Response, err error) {
	phoneVerify := utils.PhoneVerify(phone)
	if !phoneVerify {
		return &model.Response{
			Code: resultcode.Fail,
			Msg:  "phone or password verify error"}, nil
	}
	user, err := (*as.DAO).FindUserByPhone(phone)
	if err != nil {
		return &model.Response{Code: resultcode.Fail, Msg: "用户不存在。"}, err
	}
	return &model.Response{
		Code: resultcode.Success,
		Data: &user,
		Msg:  "success"}, nil
}
