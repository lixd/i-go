package tbk

import (
	"errors"
	"github.com/spf13/viper"
	"time"
)

var TbkConf *tbkConf

type tbkConf struct {
	AppKey    string
	AppSecret string
	Router    string
	Session   string
	Timeout   time.Duration
}

func ParseConf() (*tbkConf, error) {
	var c tbkConf
	if err := viper.UnmarshalKey("tbk", &c); err != nil {
		return &tbkConf{}, err
	}
	if c.AppSecret == "" {
		return &tbkConf{}, errors.New("tbk conf nil")
	}
	TbkConf = &c
	return &c, nil
}
