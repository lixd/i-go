package pay

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

var (
	cli           *core.Client
	wxpayConf     wxpay
	NotifyHandler *notify.Handler // 用于处理回调请求
)

type wxpay struct {
	MchID                      string `json:"mchID"`                      // 商户号 例如：15938xxxx 获取途径: 【账户中心】-->【商户信息】-->[微信支付商户号]
	MchCertificateSerialNumber string `json:"mchCertificateSerialNumber"` // 商户证书序列号 40位字符 获取途径: 【账户中心】-->【API安全】-->【管理证书】-->[证书序列号]
	MchAPIv3Key                string `json:"mchAPIv3Key"`                // 商户APIv3密钥  32位字符 获取途径: 【账户中心】-->【API安全】-->[设置APIv3密钥]
	MchPrivateKey              string `json:"mchPrivateKey"`              // 商户私钥证书路径 例如：apiclient_key.pem 获取途径: 【账户中心】-->【API安全】-->【管理证书】-->[申请证书]

	AppId     string `json:"appId"`     // AppID,获取途径: 【产品中心】-->【AppID账号管理】
	NotifyURL string `json:"notifyURL"` // 通知URL
}

func InitWxPay() {
	if err := viper.UnmarshalKey("wxpay", &wxpayConf); err != nil {
		panic(err)
	}
	log.Printf("conf:%v\n", wxpayConf)
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(wxpayConf.MchPrivateKey)
	if err != nil {
		log.Fatalf("load merchant private key error:%s", err)
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(wxpayConf.MchID, wxpayConf.MchCertificateSerialNumber, mchPrivateKey, wxpayConf.MchAPIv3Key),
	}
	cli, err = core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}
	// 获取平台证书访问器
	certVisitor := downloader.MgrInstance().GetCertificateVisitor(wxpayConf.MchID)
	NotifyHandler = notify.NewNotifyHandler(wxpayConf.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))
	log.Println("wxpay 初始化完成")
}

const (
	Desc = "xxx充值服务"
)

// PC 微信支付 PC-Native支付，生成二维码，扫码支付，返回值为二维码链接
func PC(orderId string, amount int64) (string, error) {
	amount = 1 // TODO 测试金额

	svc := native.NativeApiService{Client: cli}
	req := native.PrepayRequest{
		Appid:       core.String(conf.AppId),
		Mchid:       core.String(wxpayConf.MchID),
		Description: core.String(Desc),
		OutTradeNo:  core.String(orderId),
		TimeExpire:  core.Time(time.Now().Add(time.Minute * 30)), // 30分钟后过期
		NotifyUrl:   core.String(wxpayConf.NotifyURL),            // 必须为直接可访问的URL，不允许携带查询串，要求必须为https地址。
		Amount: &native.Amount{
			Currency: core.String("CNY"), // 默认为人民币
			Total:    core.Int64(amount), // 单位：分
		},
	}
	resp, result, err := svc.Prepay(context.Background(), req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "微信支付-PC"}).Error(err)
		return "", err
	}
	if result.Response.StatusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"scene": "微信支付-PC", "status": result.Response.StatusCode, "resp": resp}).Error(err)
		return *resp.CodeUrl, errors.New("创建订单失败")
	}

	return *resp.CodeUrl, nil
}

// H5 微信支付 H5-API 独立与Native、JSAPI、APP 的一个API，实际为中间页URL，用户访问该URL会唤醒微信进行支付
// 访问时提示：商家参数格式有误，请联系商家解决，一般为 referer 和微信平台配置值不一致导致，直接访问该URL时referer为空，必定会出现该问题。
func H5(orderId, clientIp string, amount int64) (string, error) {
	log.Printf("orderId:%s ip:%s amount:%v\n", orderId, clientIp, amount)
	amount = 1 // TODO 测试金额
	svc := h5.H5ApiService{Client: cli}
	req := h5.PrepayRequest{
		Appid:       core.String(conf.AppId),
		Mchid:       core.String(wxpayConf.MchID),
		Description: core.String(Desc),
		OutTradeNo:  core.String(orderId),
		TimeExpire:  core.Time(time.Now().Add(time.Minute * 30)),
		NotifyUrl:   core.String(wxpayConf.NotifyURL),
		Amount: &h5.Amount{
			Currency: core.String("CNY"),
			Total:    core.Int64(amount),
		},
		SceneInfo: &h5.SceneInfo{
			PayerClientIp: core.String(clientIp),
			H5Info: &h5.H5Info{
				Type: core.String("Wap"), // 示例值：iOS, Android, Wap
			},
		},
	}
	resp, result, err := svc.Prepay(context.Background(), req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "微信支付-H5"}).Error(err)
		return "", err
	}
	if result.Response.StatusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"scene": "微信支付-PC", "status": result.Response.StatusCode, "resp": resp}).Error(err)
		return *resp.H5Url, errors.New("创建订单失败")
	}
	return *resp.H5Url, nil
}

func wxpayCallback(req *http.Request) bool {
	transaction := new(payments.Transaction)
	notifyReq, err := NotifyHandler.ParseNotifyRequest(context.Background(), req, transaction)
	// 如果验签未通过，或者解密失败
	if err != nil {
		logrus.WithFields(logrus.Fields{"scenes": "支付回调-微信-验签or解密"}).Error(err)
		return false
	}
	log.Printf("订单数据: trx:%+v req:%+v\n", transaction, notifyReq)
	log.Printf("订单号:%s 状态:%s\n", *transaction.OutTradeNo, notifyReq.Summary)
	orderId := *transaction.OutTradeNo
	orderNo := *transaction.TransactionId
	fmt.Print("商户订单号:%s 微信支付订单号:%s \n", orderId, orderNo)

	// 判断是否重复
	// 	具体逻辑
	return true
}

type WxNotify struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// WechatPayCallback 微信支付回调 controller https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_4_5.shtml
func WechatPayCallback(c *gin.Context) {
	log.Println("微信支付回调")
	ok := wxpayCallback(c.Request)
	// 支付通知http应答码为200或204才会当作正常接收，当回调处理异常时，应答的HTTP状态码应为500，或者4xx。
	if ok {
		c.JSON(http.StatusOK, WxNotify{Code: "SUCCESS", Message: "成功"})
		return
	}
	c.JSON(http.StatusInternalServerError, WxNotify{Code: "FAILURE", Message: "失败"})
}
