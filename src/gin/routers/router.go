package routers

import (
	"github.com/gin-gonic/gin"
	user2 "i-go/src/gin/user"
)

func RegisterRoutes(router *gin.Engine) {

	v1Controller := &user2.V1Controller{}
	// 加载静态资源
	// router.LoadHTMLFiles("E:/illusory/MyProject/i-go/gin/templates/index.html")
	router.LoadHTMLGlob("E:/illusory/MyProject/i-go/src/gin/templates/*")
	// 设定请求url不存在的返回值
	router.NoRoute(v1Controller.NoResponse)
	// 重定向
	router.GET("/redirect", v1Controller.RedirectHandler)
	// 路由重定向
	router.GET("/routeredirect", func(c *gin.Context) {
		c.Request.URL.Path = "/redirect"
		router.HandleContext(c)
	})

	// route 分组
	v1 := router.Group("/v1")
	// 使用中间件
	// v1.Use(middleware.AuthMiddleWare())
	{
		v1.GET("/", v1Controller.IndexHandler)
		v1.GET("/ping", v1Controller.PingHandler)
		v1.GET("/user", v1Controller.UserHandler)
		v1.GET("/index", v1Controller.LoadHTMLGlobHandler)
		v1.GET("/JSONP?callback=x", v1Controller.JSONPHandler)
		// 表单绑定
		v1.POST("/login", v1Controller.LoginHandler)
		// 快速参数匹配
		v1.GET("/login2/:name/:password", v1Controller.Login2Handler)
		// 普通参数匹配 http://localhost:8080/v1/login3?name=root&password=root
		v1.GET("/login3", v1Controller.Login3Handler)
		v1.GET("/login4", v1Controller.Login4Handler)
		// 文件上传
		v1.POST("/upload", v1Controller.UploadHandler)
		v1.GET("/someDataFromReader", v1Controller.SomeDataFromReaderHandler)
		v1.GET("/long_async", v1Controller.LongAsyncHandler)
		v1.GET("/long_sync", v1Controller.LongSyncHandler)
		v1.GET("/login5", v1Controller.RequestBodyHandler)
		v1.POST("/login6", v1Controller.QueryMap)
		v1.POST("/checkbox", v1Controller.CheckBoxHandler)
		v1.POST("/paramsbind", v1Controller.ParamsBind)
		v1.GET("/paramsshouldbind", v1Controller.ParamsShouldBind)
		v1.GET("/cookie", v1Controller.CookieHandler)
	}

	v2Controller := &user2.ParamBind{}
	v2 := router.Group("/v2")
	{
		v2.GET("/upload", v2Controller.UploadFileHandler)
		// URL 路径参数匹配
		v2.GET("/param/:user/:password", v2Controller.ParamHandler)
		// URL 查询参数参数匹配
		v2.GET("/query", v2Controller.QueryHandler)
		v2.POST("/postform", v2Controller.PostFormHandler)
		v2.GET("/querymap", v2Controller.QueryMapHandler)
		v2.POST("/postformmap", v2Controller.PostFormMapHandler)
		// 文件上传
		v2.POST("/formfile", v2Controller.FormFileHandler)
		v2.POST("/multipartform", v2Controller.MultipartFormHandler)
	}

	v3Controller := &user2.ModelBind{}
	v3 := router.Group("/v3")
	{
		v3.GET("/bind", v3Controller.BindHandler)
		v3.GET("/bindjson", v3Controller.BindJSONHandler)
		v3.GET("/shouldbind", v3Controller.ShouldBindHandler)
		v3.GET("/shouldbindjson", v3Controller.ShouldBindJSONHandler)
	}
	v4Controller := user2.V4Controller{}
	v4 := router.Group("/v4")
	{
		v4.GET("/string", v4Controller.StringHandler)
		v4.GET("/json", v4Controller.JSONHandler)
		v4.GET("/data", v4Controller.DATAHandler)
		v4.GET("/html", v4Controller.HTMLHandler)
	}

}
