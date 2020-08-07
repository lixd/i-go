package logic

import (
	"github.com/lixd/vaptcha-sdk-go"
	"github.com/lixd/vaptcha-sdk-go/examples/constant"
	"github.com/lixd/vaptcha-sdk-go/model"
	"net/http"
)

// Offline VAPTCHA离线验证接口
func Offline(writer http.ResponseWriter, request *http.Request) {
	req := queryParams(request)

	option := func(options *vaptcha.Options) {
		// options.Vid = constant.Vid
		options.Vid = "offline" // 如果是测试离线模式 则vid直接传offline即可
		options.SecretKey = constant.SecretKey
		options.Scene = constant.Scene
	}
	v := vaptcha.NewVaptcha(option)
	// invoke sdk offline
	result := v.Offline(req)

	_, _ = writer.Write([]byte(result))
}

// queryParams parse query params
func queryParams(request *http.Request) (ret model.Offline) {
	query := request.URL.Query()
	actions := query["offline_action"]
	callbacks := query["callback"]
	knocks := query["knock"]
	userCodes := query["v"]
	if len(actions) != 0 {
		ret.Action = actions[0]
	}
	if len(callbacks) != 0 {
		ret.Callback = callbacks[0]
	}
	if len(knocks) != 0 {
		ret.Knock = knocks[0]
	}
	if len(userCodes) != 0 {
		ret.UserCode = userCodes[0]
	}
	return
}
