package user

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type Simple struct {
}

// ReturnHTML 返回整个页面
func (Simple *Simple) ReturnHTML(c *gin.Context) {
	// 替换掉要变的那部分代码
	htmlNew := strings.ReplaceAll(html, "{{code}}", code)
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, htmlNew)
}

const pathAdServer = ""

// ReturnHTMLTransform 从旧服务器转发到新的这边来
func (Simple *Simple) ReturnHTMLTransform(c *gin.Context) {
	resp, err := http.Get(pathAdServer)
	if err != nil {
		c.String(200, err.Error())
		return
	}
	_, _ = io.Copy(c.Writer, resp.Body)
	resp.Body.Close()
}

const code = `<script type="text/javascript">var jd_union_unid="1002590618",jd_ad_ids="508:6",jd_union_pid="CILlifqTLhCao4neAxoAII2r5JYLKgA=";var jd_width=0;var jd_height=0;var jd_union_euid="";var p="ABMGVB9cEQAQA2VEH0hfIlgRRgYlXVZaCCsfSlpMWGVEH0hfImMjHRxlAhNUNhpYYAZzZAlgGEZ1cFFZF2sXAxMGURxeEwUaN1UaWhYGGgZSG1IlMk1DCEZrXmwTNwpfBkgyEgNcH1MUBxEDVR9YEjITN2Ur";</script><script type="text/javascript" charset="utf-8" src="//u-x.jd.com/static/js/auto.js"></script>`

const html = `
<!DOCTYPE html>
<html lang="">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>iframe</title>
</head>

<body>
  <div>{{code}}</div>
</body>

</html>`
