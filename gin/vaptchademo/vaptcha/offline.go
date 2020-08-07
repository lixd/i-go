package vaptcha

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	. "i-go/gin/vaptchademo/vaptcha/constant"
	"i-go/gin/vaptchademo/vaptcha/model"
	"i-go/gin/vaptchademo/vaptcha/utils"
	"net/http"
	"strings"
	"time"
)

// Offline offline mode
func (v *vaptcha) Offline(req model.Offline) (ret string) {
	if req.Action == ActionGet {
		image := v.getImage()
		return buildJSONP(req.Callback, image)
	}
	if req.Action == ActionVerify {
		validate := v.offlineValidate(req.Knock, req.UserCode)
		return buildJSONP(req.Callback, validate)
	}
	return
}

// getImage generate imgid and knock
func (v *vaptcha) getImage() (ret model.Image) {
	key, offlineState := v.getKey()
	if offlineState == 0 {
		ret.Code = Fail
		ret.Msg = "VAPTCHA is not in offline mode"
		return
	}
	if key == "" {
		ret.Code = Fail
		ret.Msg = "get offline key fail"
		return
	}
	imgId := v.generateImgId(key)
	knock := v.generateKnock(imgId)
	ret.Code = Success
	ret.ImgId = imgId
	ret.Knock = knock
	return
}

// offlineValidate validate fronted path in offline mode
func (v *vaptcha) offlineValidate(knock, userCode string) (ret model.Validate) {
	key, _ := v.getKey()
	if key == "" {
		ret.Code = Fail
		ret.Msg = "get offline key fail"
		return
	}
	verifyKey := v.GetVerifyKey(knock, userCode)
	url := fmt.Sprintf("%s/%s/%s", v.options.OfflineVerifyURL, key, verifyKey)
	_, code := utils.HttpGet(url)
	if code != http.StatusOK {
		ret.Code = Fail
		ret.Msg = "http fail"
		return
	}
	token := v.generateToken(knock)
	ret.Code = Success
	ret.Token = OfflineMode + knock + token
	return
}

// buildJSONP
// e.g. VaptchaJsonp1596787212811({code: "0103", imgid: "5eea673b825aa0a38c457d68181d86c1", knock: "8fbcf455da9c444ab384e1b3a1d07872"})
func buildJSONP(callBack string, data interface{}) (jsonp string) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	jsonp = callBack + "(" + string(bytes) + ")"
	return
}

// GetVerifyKey
func (v *vaptcha) GetVerifyKey(knock string, userCode string) string {
	// kconk-->unix(10)+imgId(32)
	kconk := v.options.Cache.Get(knock)
	// Delete the cache to ensure that each picture can only be verified once
	v.options.Cache.Delete(knock)
	// verifyKey-->MD5(userCode+imgId)
	imgId := kconk[10:]
	verifyKey := utils.MD5(userCode + imgId)
	return verifyKey
}

func (v *vaptcha) generateImgId(key string) (imgId string) {
	randStr := utils.RandStr()
	str := key + randStr
	imgId = utils.MD5(str)
	fmt.Printf("key:%v randStr:%v imgId:%v \n", key, randStr, imgId)
	return
}
func (v *vaptcha) generateKnock(imgId string) (knock string) {
	randKnock := uuid.NewV4().String()
	knock = strings.ReplaceAll(randKnock, "-", "")
	v.options.Cache.Set(knock, fmt.Sprintf("%v%s", time.Now().Unix(), imgId))
	return
}
func (v *vaptcha) generateToken(knock string) string {
	randStr := uuid.NewV4().String()
	token := strings.ReplaceAll(randStr, "-", "")
	v.options.Cache.Set(knock, fmt.Sprintf("%v%s", time.Now().Unix(), token))
	return token
}

// 获取offlineKey
func (v *vaptcha) getKey() (string, int) {
	var (
		bt []byte
	)
	url := v.options.ChannelURL + "/" + v.options.Vid
	bt, _ = utils.HttpGet(url)

	var ol = struct {
		Key   string `json:"offline_key"`
		State int    `json:"offline_state"`
	}{}
	_ = json.Unmarshal(bt, &ol)
	return ol.Key, ol.State
}
