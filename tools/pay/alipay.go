package pay

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
)



var (
	AliClient *alipay.Client
	conf      ConfAlipay
)

type ConfAlipay struct {
	IsProduction     bool   `json:"isProduction"`
	AppId            string `json:"appId"`
	PrivateKey       string `json:"privateKey"`
	AlipayRootCert   string `json:"alipayRootCert"`
	AliPayPublicCert string `json:"alipayPublicCert"`
	AppPublicKeyCert string `json:"appPublicKeyCert"`
	NotifyUrl        string `json:"notifyUrl"`
	ReturnUrl        string `json:"returnUrl"`
	ReturnUrlH5      string `json:"returnUrlH5"`
}

func (ca ConfAlipay) String() string {
	bt, _ := json.Marshal(ca)
	return string(bt)
}

const (
	TradeSuccess = alipay.TradeStatusSuccess
)

func Init() {
	if err := viper.UnmarshalKey("alipay", &conf); err != nil {
		panic(err)
	}
	if conf.AppId == "" {
		panic("未获取到alipay配置文件")
	}
	var err error
	AliClient, err = alipay.New(conf.AppId, conf.PrivateKey, conf.IsProduction)
	if err != nil {
		panic(err)
	}
	err = AliClient.LoadAppPublicCertFromFile(conf.AppPublicKeyCert) // 加载应用公钥证书
	if err != nil {
		panic(err)
	}
	err = AliClient.LoadAliPayRootCertFromFile(conf.AlipayRootCert) // 加载支付宝根证书
	if err != nil {
		panic(err)
	}
	err = AliClient.LoadAliPayPublicCertFromFile(conf.AliPayPublicCert) // 加载支付宝公匙证书
	if err != nil {
		panic(err)
	}
	log.Println("alipay 初始化完成")
}

// UnifiedOrder PC支付
func UnifiedOrder(outTradeNo, passBack string, amount float64) (string, error) {
	var p alipay.TradePagePay
	p.NotifyURL = conf.NotifyUrl
	p.ReturnURL = conf.ReturnUrl
	p.Subject = "xxx充值服务"
	p.OutTradeNo = outTradeNo
	p.TotalAmount = strconv.FormatFloat(amount, 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	p.PassbackParams = passBack
	urlObj, err := AliClient.TradePagePay(p)
	if err != nil {
		return "", err
	}
	return urlObj.String(), nil
}

// WapPay H5支付
func WapPay(outTradeNo, passBack string, amount float64) (string, error) {
	var p alipay.TradeWapPay
	p.NotifyURL = conf.NotifyUrl
	p.ReturnURL = conf.ReturnUrlH5
	p.Subject = "xxx充值服务"
	p.OutTradeNo = outTradeNo
	p.TotalAmount = strconv.FormatFloat(amount, 'f', 2, 64)
	// p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	p.ProductCode = "QUICK_WAP_WAY"
	p.PassbackParams = passBack
	urlObj, err := AliClient.TradeWapPay(p)
	if err != nil {
		return "", err
	}
	return urlObj.String(), nil
}

func AliTrans(amount, fee float64, account, realName, withdrawNo string) (bool, string, error) {
	payeeInfo := &alipay.PayeeInfo{
		Identity:     account,
		IdentityType: "ALIPAY_USER_ID",
		Name:         realName,
	}

	req := alipay.FundTransUniTransfer{}
	req.TransAmount = fmt.Sprintf("%v", amount)
	req.OutBizNo = withdrawNo
	req.ProductCode = "TRANS_ACCOUNT_NO_PWD"
	req.PayeeInfo = payeeInfo
	req.OrderTitle = "xxx提现服务"
	req.BizScene = "DIRECT_TRANSFER"
	req.Remark = fmt.Sprintf("含%v手续费", fee)
	result, err := AliClient.FundTransUniTransfer(req)
	if err != nil {
		return false, "", err
	}
	return result.IsSuccess(), result.Content.SubCode, nil
}

func alipayCallback(req *http.Request) bool {
	var (
		outTradeNo     string
		orderNo        string
	)
	notify, err := AliClient.GetTradeNotification(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "支付-回调-alipay"}).Error(err)
		return false
	}
	if notify.TradeStatus == TradeSuccess {
		logrus.WithFields(logrus.Fields{"Scenes": "支付-回调-alipay-返回值不为success"}).Error(err)
		return false
	}
	// 支付结果
	if notify.TradeStatus == TradeSuccess {
		outTradeNo = notify.OutTradeNo
		orderNo = notify.TradeNo
	}
	fmt.Printf("商户订单号:%s 支付宝订单号:%s \n",outTradeNo,orderNo)
	// 	订单重复校验
	// 	逻辑处理
	return true
}

// AlipayCallback 支付宝支付回调 controller 层
func AlipayCallback(c *gin.Context) {
	// 收到异步通知后，商家输出success是表示消息获取成功，支付宝就会停止发送异步，如果输出fail，表示消息获取失败，支付宝会重新发送消息到异步地址。
	ok := alipayCallback(c.Request)
	if !ok {
		c.String(http.StatusOK, "fail")
		return
	}
	c.String(http.StatusOK, "success")
}