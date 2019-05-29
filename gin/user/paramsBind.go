package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ParamBind struct {
}

// c.Param() URL路径参数 /param/admin/root
func (ParamBind *ParamBind) ParamHandler(c *gin.Context) {
	paramUser := c.Param("user")         // user=admin
	paramPassword := c.Param("password") // password=root
	log.Printf("c.Param()  user=%v password=%v", paramUser, paramPassword)
}

// c.Query() 查询参数 /query?user=admin&password=root
// c.DefaultQuery()
func (ParamBind *ParamBind) QueryHandler(c *gin.Context) {
	queryUser := c.Query("user")         // user=admin
	queryPassword := c.Query("password") // password=root
	// queryPassword:=c.DefaultQuery("user","")
	log.Printf("c.Query()  user=%v password=%v", queryUser, queryPassword)
}

// c.PostForm() POST 表单
func (ParamBind *ParamBind) PostFormHandler(c *gin.Context) {
	PostFormUser := c.PostForm("user")         // user=admin
	PostFormPassword := c.PostForm("password") // password=root
	// PostFormPassword := c.DefaultPostForm("password","none") // 不存在时有默认值"none"
	log.Printf("c.PostForm()  user=%v password=%v", PostFormUser, PostFormPassword)
}

// c.QueryMap() 返回值为 map[string]string
// /querymap?user[1]=admin&password[1]=root&user[2]=admin2&password[2]=root2
func (ParamBind *ParamBind) QueryMapHandler(c *gin.Context) {
	QueryMapUsers := c.QueryMap("user")    // user=admin
	QueryMapPwds := c.QueryMap("password") // password=root
	log.Printf("c.QueryMap()  user=%v password=%v", QueryMapUsers, QueryMapPwds)
	// QueryArrUsers := c.QueryArray("user")
	// QueryArrPwds := c.QueryArray("password")
	// log.Printf("c.QueryArray()  user=%v password=%v", QueryArrUsers, QueryArrPwds)
}

// c.PostForm() POST 表单 返回值 map[string]string
func (ParamBind *ParamBind) PostFormMapHandler(c *gin.Context) {
	postFormMapUsers := c.PostFormMap("user")
	postFormMapPassword := c.PostFormMap("password")
	log.Printf("c.QueryMap()  user=%v password=%v", postFormMapUsers, postFormMapPassword)
}
func (ParamBind *ParamBind) UploadFileHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

// 单文件上传  根据 key 获取form表单的第一个文件
func (ParamBind *ParamBind) FormFileHandler(c *gin.Context) {
	// FormFile returns the first file for the provided form key.
	header, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		log.Printf("c.FormFile() error error=%v", err)
		return
	}
	log.Printf("c.FormFile()  file name=%v error=%v", header.Filename, err)
}

// 多文件上传
// 返回类型如下
/*
	type Form struct {
		Value map[string][]string
		File  map[string][]*FileHeader
	}
*/

func (ParamBind *ParamBind) MultipartFormHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	log.Printf("c.MultipartForm() form=%v err=%v", form, err)
	headers := form.File
	for value, file := range headers {
		log.Printf("c.MultipartForm()  form.File.value=%v form.File.file=%v", value, file)
	}
}
