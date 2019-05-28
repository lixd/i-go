package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"time"
)

func main() {
	// 修改模式
	// gin.SetMode(gin.ReleaseMode)
	// 记录日志到文件
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }
	// Default方法创建一个路由handler。
	router := gin.Default()
	// 加载静态资源
	// router.LoadHTMLFiles("E:/illusory/MyProject/i-go/gin/templates/index.html")
	router.LoadHTMLGlob("E:/illusory/MyProject/i-go/gin/templates/*")
	// 设定请求url不存在的返回值
	router.NoRoute(NoResponse)
	// 分组
	v1 := router.Group("/v1")
	{
		v1.GET("/", IndexHandler)
		v1.GET("/ping", PingHandler)
		v1.GET("/user", UserHandler)
		v1.GET("/index", LoadHTMLGlobHandler)
		v1.GET("/JSONP?callback=x", JSONPHandler)
		// 表单绑定
		v1.POST("/login", LoginHandler)
		// 快速参数匹配
		v1.GET("/login2/:name/:password", Login2Handler)
		// 普通参数匹配 http://localhost:8080/v1/login3?name=root&password=root
		v1.GET("/login3", Login3Handler)
		v1.GET("/login4", Login4Handler)
		// 文件上传
		v1.POST("/upload", UploadHandler)
		v1.GET("/someDataFromReader", SomeDataFromReaderHandler)
		v1.GET("/long_async", LongAsyncHandler)
		v1.GET("/long_sync", LongSyncHandler)
		v1.GET("/login5", RequestBodyHandler)
		v1.POST("/login6", QueryMap)
		v1.POST("/checkbox", CheckBoxHandler)
	}

	router.Run(":8080")
}
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func CheckBoxHandler(c *gin.Context) {
	var myForm myForm
	c.ShouldBind(&myForm)
	c.JSON(http.StatusOK, gin.H{"colors": myForm.Colors})
}
func QueryMap(c *gin.Context) {
	// map
	ids := c.QueryMap("ids")
	names := c.QueryMap("names")
	c.String(http.StatusOK, "ids: %v; names: %v", ids, names)
}
func RequestBodyHandler(c *gin.Context) {
	user := LoginForm{}
	class := Class{}
	if c.ShouldBindBodyWith(&user, binding.JSON) == nil {
		c.String(http.StatusOK, "the body should be LoginForm  user=%v password= %v", user.User, user.Password)
	} else if c.ShouldBindBodyWith(&class, binding.JSON) == nil {
		c.String(http.StatusOK, "he body should be class id=%v number= %v", class.Id, class.Number)
	}
}
func LongSyncHandler(c *gin.Context) {
	time.Sleep(time.Second * 5)
	// 未使用goroutine 不用创建副本
	log.Println("Done in path:" + c.Request.URL.Path)
}
func LongAsyncHandler(c *gin.Context) {
	// 创建在 goroutine 中使用的副本
	cCp := c.Copy()
	go func() {
		time.Sleep(time.Second * 5)
		log.Println("Done in path:" + cCp.Request.URL.Path)
	}()
}
func SomeDataFromReaderHandler(c *gin.Context) {
	response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}
func UploadHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
func NoResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "page not found"})
}

// hello world
func PingHandler(c *gin.Context) {
	// gin.H 是 map[string]interface{} 的一种快捷方式
	c.JSON(200, gin.H{"message": "pong"})
}

// AsciiJSON
func UserHandler(c *gin.Context) {
	User1 := User{"illusory", 23, "CQ"}
	c.AsciiJSON(http.StatusOK, User1)
}

// HTML 渲染
func LoadHTMLGlobHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{"title": "Main website"})
}

func PusherHandler(c *gin.Context) {
	if pusher := c.Writer.Pusher(); pusher != nil {
		if err := pusher.Push("/assets/app.js", nil); err != nil {
			log.Printf("Push failed err= %v \n", err)
		}
		c.HTML(http.StatusOK, "https", gin.H{"status": "success"})
	}
}

func JSONPHandler(c *gin.Context) {
	User1 := User{"illusory", 23, "CQ"}
	c.JSONP(http.StatusOK, User1)
}

// form 表单绑定
func LoginHandler(c *gin.Context) {
	var form LoginForm
	if c.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}
}

// 快速参数匹配
func Login2Handler(c *gin.Context) {
	name := c.Param("name")
	password := c.Param("password")
	id, err := c.GetQuery("id")
	if !err {
		name = "the key is not exist"
	}
	c.String(http.StatusOK, "name=%v password=%v id=%v ", name, password, id)
}

// 普通参数匹配
func Login3Handler(c *gin.Context) {
	// 找不到时设置默认值 admin
	name := c.DefaultQuery("name", "admin")
	// 找不到时直接为空
	password := c.Query("password")
	c.String(http.StatusOK, "name=%v password=%v ", name, password)
}

// 普通参数匹配
func Login4Handler(c *gin.Context) {
	var form LoginForm
	if c.ShouldBindQuery(&form) == nil {
		name := form.User
		password := form.Password
		c.String(http.StatusOK, "name=%v password=%v ", name, password)
	}
}

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
type LoginForm struct {
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
}
type Class struct {
	Id     string `json:"id" form:"id"`
	Number int    `json:"number" form:"number"`
}
type myForm struct {
	Colors []string `form:"colors[]"`
}
