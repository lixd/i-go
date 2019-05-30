package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"time"
)

type V1Controller struct {
}

func (v1Controller *V1Controller) IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func (v1Controller *V1Controller) CheckBoxHandler(c *gin.Context) {
	var myForm myForm
	c.ShouldBind(&myForm)
	c.JSON(http.StatusOK, gin.H{"colors": myForm.Colors})
}
func (v1Controller *V1Controller) QueryMap(c *gin.Context) {
	// map
	ids := c.QueryMap("ids")
	names := c.QueryMap("names")
	c.String(http.StatusOK, "ids: %v; names: %v", ids, names)
}
func (v1Controller *V1Controller) RequestBodyHandler(c *gin.Context) {
	user := LoginForm{}
	class := Class{}
	if c.ShouldBindBodyWith(&user, binding.JSON) == nil {
		c.String(http.StatusOK, "the body should be LoginForm  user=%v password= %v", user.User, user.Password)
	} else if c.ShouldBindBodyWith(&class, binding.JSON) == nil {
		c.String(http.StatusOK, "he body should be class id=%v number= %v", class.Id, class.Number)
	}
}
func (v1Controller *V1Controller) LongSyncHandler(c *gin.Context) {
	time.Sleep(time.Second * 5)
	// 未使用goroutine 不用创建副本
	log.Println("Done in path:" + c.Request.URL.Path)
}
func (v1Controller *V1Controller) LongAsyncHandler(c *gin.Context) {
	// 创建在 goroutine 中使用的副本
	cCp := c.Copy()
	go func() {
		time.Sleep(time.Second * 5)
		log.Println("Done in path:" + cCp.Request.URL.Path)
	}()
}
func (v1Controller *V1Controller) SomeDataFromReaderHandler(c *gin.Context) {
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
func (v1Controller *V1Controller) UploadHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
func (v1Controller *V1Controller) NoResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "page not found"})
}

// hello world
func (v1Controller *V1Controller) PingHandler(c *gin.Context) {
	// gin.H 是 map[string]interface{} 的一种快捷方式
	c.JSON(200, gin.H{"message": "pong"})
}

// AsciiJSON
func (v1Controller *V1Controller) UserHandler(c *gin.Context) {
	User1 := User{"illusory", 23, "CQ"}
	c.AsciiJSON(http.StatusOK, User1)
}

// HTML 渲染
func (v1Controller *V1Controller) LoadHTMLGlobHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{"title": "Main website"})
}

func (v1Controller *V1Controller) PusherHandler(c *gin.Context) {
	if pusher := c.Writer.Pusher(); pusher != nil {
		if err := pusher.Push("/assets/app.js", nil); err != nil {
			log.Printf("Push failed err= %v \n", err)
		}
		c.HTML(http.StatusOK, "https", gin.H{"status": "success"})
	}
}

func (v1Controller *V1Controller) JSONPHandler(c *gin.Context) {
	User1 := User{"illusory", 23, "CQ"}
	c.JSONP(http.StatusOK, User1)
}

// form 表单绑定
func (v1Controller *V1Controller) LoginHandler(c *gin.Context) {
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
func (v1Controller *V1Controller) Login2Handler(c *gin.Context) {
	name := c.Param("name")
	password := c.Param("password")
	id, err := c.GetQuery("id")
	if !err {
		name = "the key is not exist"
	}
	c.String(http.StatusOK, "name=%v password=%v id=%v ", name, password, id)
}

// 普通参数匹配
func (v1Controller *V1Controller) Login3Handler(c *gin.Context) {
	// 找不到时设置默认值 admin
	name := c.DefaultQuery("name", "admin")
	// 找不到时直接为空
	password := c.Query("password")
	c.String(http.StatusOK, "name=%v password=%v ", name, password)
}

// 普通参数匹配
func (v1Controller *V1Controller) Login4Handler(c *gin.Context) {
	var form LoginForm
	if c.ShouldBindQuery(&form) == nil {
		name := form.User
		password := form.Password
		c.String(http.StatusOK, "name=%v password=%v ", name, password)
	}
}
func (v1Controller *V1Controller) ParamsBind(c *gin.Context) {
	var form LoginForm
	err := c.Bind(&form)
	log.Printf("ShouldBind error= %v  form.user=%v form.password=%v", err, form.User, form.Password)

}
func (v1Controller *V1Controller) ParamsShouldBind(c *gin.Context) {
	var form LoginForm
	err := c.ShouldBind(&form)
	log.Printf("ShouldBind error= %v  form.user=%v form.password=%v", err, form.User, form.Password)
}

func (v1Controller *V1Controller) CookieHandler(c *gin.Context) {
	cookie, e := c.Cookie("gin_cookie")
	if e != nil {
		cookie = "NoSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
}
func (v1Controller *V1Controller) RedirectHandler(c *gin.Context) {
	log.Print("HTTP重定向。。")
	c.Redirect(http.StatusMovedPermanently, "https://www.lixueduan.com")
}
func (v1Controller *V1Controller) HandleContextHandler(c *gin.Context, r *gin.Engine) {
	log.Print("路由重定向。。")
	c.Request.URL.Path = "/redirect"
	r.HandleContext(c)
}
