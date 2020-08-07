package vaptchasdk_back

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	char        = "0123456789abcdef"
	baseUrl     = "https://channel2.vaptcha.com"
	verifyUrl   = "https://offline.vaptcha.com"
	offlineMode = "offline"
	// 手势验证
	Success = "0103"
	Fail    = "0104"
	// 二次验证
	VerifySuccess   = 1
	SecondVerifyUrl = "http://0.wlinno.com/verify"
	// 验证单元信息 自行替换
	SecretKey = "035b20d63ead49269a5e1644976a6a2e" // 验证单元key
	Vid       = "5dba48a598bb11d350b45728"         // 验证单元id
	Scene     = "0"                                // 场景值
)

var (
	once    sync.Once
	Vaptcha *vaptcha
	mem     sync.Map
)

type vaptcha struct {
	cache   storageMedium
	baseUrl string
	// 验证单元信息
	Vid       string
	SecretKey string
	Scene     string
}

func New(vid, secretKey, scene string) *vaptcha {
	once.Do(func() {
		Vaptcha = &vaptcha{
			baseUrl:   baseUrl,
			Vid:       vid,
			SecretKey: secretKey,
			Scene:     scene}
	})
	return Vaptcha
}

// -------------------public-------------------

func (v *vaptcha) Verify(token string, ip string) offlineResp {
	// 根据token的前缀判断是正常模式还是离线模式的二次验证
	if len(token) < 7 {
		panic("token错误")
	}
	if token[:7] == offlineMode {
		// 离线模式二次验证
		result := v.offlineVerify(token)
		var resp offlineResp
		err := json.Unmarshal([]byte(result), &resp)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
		return resp
	} else {
		// 正常模式二次验证
		// 初始化接口参数
		res, err := http.PostForm(SecondVerifyUrl, url.Values{"id": {v.Vid}, "secretkey": {v.SecretKey},
			"scene": {v.Scene}, "ip": {ip}, "token": {token}})
		if err != nil {
			panic(err)
		}
		// 将响应体转换为 byte slice
		bt, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		var resp offlineResp
		err = json.Unmarshal(bt, &resp)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp)
		return resp
	}
}

// 离线模式交互接口
func (v *vaptcha) Offline(action, callBack, vid, knock, userCode string) string {
	var result string
	if action == "get" {
		image := v.getImage(vid)
		result = fmt.Sprintf("%s%s%s%s", callBack, "(", image, ")")
	} else {
		res := v.validate(knock, userCode, vid)
		result = fmt.Sprintf("%s%s%s%s", callBack, "(", res, ")")
	}
	return result
}

type offlineResp struct {
	Msg     string `json:"msg"`
	Success int    `json:"success"`
	Score   int    `json:"score"`
}

// ---------------------------private-----------------------------

// 离线模式二次验证
func (v *vaptcha) offlineVerify(token string) string {
	m := make(map[string]interface{})
	m["msg"] = "验证失败"

	// token=offline(7)+knock(32)+uuid(32)
	knock := token[7:39]
	uid := token[39:]
	// cacheToken:unix(10)+uuid(32)
	cacheToken := v.cache.Get(knock)
	realToken := cacheToken[10:]
	if realToken == uid {
		// 验证成功则移除token
		v.cache.Delete(knock)

		m["success"] = VerifySuccess
		m["msg"] = "验证通过"
		m["score"] = 100
	}
	return v.mapToJson(m)
}

// 获取离线模式验证图
func (v *vaptcha) getImage(vid string) string {
	var m = make(map[string]interface{})
	key, offlineState := v.getKey(vid)
	if offlineState == 0 {
		m["code"] = Fail
		m["msg"] = "VAPTCHA未进入离线模式"
		return v.mapToJson(m)
	}
	if key == "" {
		m["code"] = Fail
		m["msg"] = "离线key获取失败"
		return v.mapToJson(m)
	}
	randStr := v.getRandomStr()
	str := key + randStr
	imgId := v.md5Encode(str)
	fmt.Printf("key:%v randStr:%v imgId:%v \n", key, randStr, imgId)
	randKnock := uuid.NewV4().String()
	knock := strings.ReplaceAll(randKnock, "-", "")
	// key:knock value:unix(10)+imgId(32)
	v.cache.Set(knock, fmt.Sprintf("%v%s", time.Now().Unix(), imgId))
	m["code"] = Success
	m["imgid"] = imgId
	m["knock"] = knock
	return v.mapToJson(m)
}

// 轨迹验证
func (v *vaptcha) validate(knock, userCode, vid string) string {
	var m = make(map[string]interface{})
	key, _ := v.getKey(vid)
	if key == "" {
		m["code"] = Fail
		m["msg"] = Fail
		return v.mapToJson(m)
	}
	// imgData:unix(10)+imgId(32)
	imgData := v.cache.Get(knock)
	// 删除缓存保证每张图片只能被验证一次
	v.cache.Delete(knock)
	// verifyKey: md5(userCode+imgId)
	imgId := imgData[10:]
	verifyKey := v.md5Encode(userCode + imgId)
	_, code := v.httpRequest(fmt.Sprintf("%s/%s/%s", verifyUrl, key, verifyKey))
	if code == http.StatusOK {
		randStr := uuid.NewV4().String()
		token := strings.ReplaceAll(randStr, "-", "")
		v.cache.Set(knock, fmt.Sprintf("%v%s", time.Now().Unix(), token))
		m["code"] = Success
		m["token"] = fmt.Sprintf("%s%s%s", "offline", knock, token)
		return v.mapToJson(m)
	} else {
		m["code"] = Fail
		m["msg"] = Fail
		return v.mapToJson(m)
	}
}

func (v *vaptcha) httpRequest(url string) ([]byte, int) {
	var (
		client  *http.Client
		request *http.Request
		resp    *http.Response
		err     error
	)
	client = &http.Client{}
	request, _ = http.NewRequest(http.MethodGet, url, nil)
	resp, err = client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, http.StatusNotFound
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode
}

func (v *vaptcha) md5Encode(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func (v *vaptcha) getRandomStr() string {
	charArr := []byte(char)
	var result string
	for i := 0; i < 4; i++ {
		result = fmt.Sprintf("%s%v", result, string(charArr[rand.Intn(16)]))
	}
	return result
}

// 获取offlineKey
func (v *vaptcha) getKey(vid string) (string, int) {
	var (
		bt []byte
	)
	bt, _ = v.httpRequest(fmt.Sprintf("%s/%s/%s", v.baseUrl, "config", vid))

	var ol = struct {
		Key   string `json:"offline_key"`
		State int    `json:"offline_state"`
	}{}
	_ = json.Unmarshal(bt, &ol)
	return ol.Key, ol.State
}

func (v *vaptcha) mapToJson(m map[string]interface{}) string {
	bt, _ := json.Marshal(m)
	return string(bt)
}

// 移除过期cache
func (v *vaptcha) RemoveExpireSess() {
	for {
		f := func(key, value interface{}) bool {
			var val = fmt.Sprintf("%v", value)
			if len(val) > 10 {
				unix, _ := strconv.ParseInt(val[:10], 10, 64)
				timeSpan := time.Now().Unix() - unix
				// 大于3分钟过期
				if timeSpan > 3*60 || timeSpan < -60 {
					v.cache.Delete(key)
				}
			} else {
				// 错误key直接删除
				v.cache.Delete(key)
			}
			return true
		}
		v.cache.Range(f)
		time.Sleep(time.Second * 1)
	}
}

// 缓存操作
type storageMedium struct {
}

func (storageMedium) Get(key interface{}) string {
	if value, ok := mem.Load(key); ok {
		return value.(string)
	}
	return ""
}
func (storageMedium) Set(key, value interface{}) {
	mem.Store(key, value)
}
func (storageMedium) Delete(key interface{}) {
	mem.Delete(key)
}
func (storageMedium) Range(f func(key, value interface{}) bool) {
	mem.Range(f)
}
