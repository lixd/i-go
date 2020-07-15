package gotbk

const (
	// 测试环境
	DevHttp = "http://gw.api.tbsandbox.com/router/rest"
	// ISV软件上线之后使用的环境，此环境的入口与正式测试环境一致，只不过应用上线之后，流量限制会进行打开，具体流量限制与应用所属类目有关，比如服务市场类的应用，限制API调用为100万次/天。
	ReleaseHttp  = "http://gw.api.taobao.com/router/rest"
	ReleaseHttps = "https://eco.taobao.com/router/rest"
	// 海外环境也属于正式环境的一种，主要是给海外（欧美国家）ISV使用，对于海外的ISV，使用海外环境会比国内环境的性能高一倍。
	OverseasHttp  = "http://api.taobao.com/router/rest"
	OverseasHttps = "https://api.taobao.com/router/rest"
)
