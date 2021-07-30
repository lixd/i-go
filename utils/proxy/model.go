package proxyutil

// ProxyIP 代理IP返回数据
type ProxyIP struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []Data `json:"data"`
}

// Data 具体IP信息
type Data struct {
	IP         string `json:"ip"`
	Port       int    `json:"port"`
	ExpireTime string `json:"expire_time"`
	City       string `json:"city"`
	Isp        string `json:"isp"`
}
