package vaptcha

import (
	. "i-go/gin/vaptchademo/vaptcha/constant"
	"sync"
)

type vaptcha struct {
	options Options
}
type Options struct {
	Cache            ICache
	ChannelURL       string
	OfflineVerifyURL string
	OnlineVerifyURL  string
	// 验证单元信息
	Vid       string
	SecretKey string
	Scene     string
}

type option func(*Options)

// configure merge custom options
func configure(v *vaptcha, opts ...option) {
	for _, opt := range opts {
		opt(&v.options)
	}
}
func build(opts ...option) *vaptcha {
	// default config
	options := Options{
		Cache:            vCache{},
		ChannelURL:       ChannelURL,
		OfflineVerifyURL: OfflineVerifyURL,
		OnlineVerifyURL:  OnlineVerifyURL,
		Vid:              Vid,
		SecretKey:        SecretKey,
		Scene:            Scene,
	}
	// new cli by default config
	v := &vaptcha{
		options: options,
	}
	//	merge config
	configure(v, opts...)
	return v
}

var (
	once    sync.Once
	Vaptcha *vaptcha
)

func NewVaptcha(opts ...option) *vaptcha {
	once.Do(func() {
		Vaptcha = build(opts...)
		go Vaptcha.RemoveExpireKey()
	})
	return Vaptcha
}
