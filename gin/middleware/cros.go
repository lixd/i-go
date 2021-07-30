package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 处理跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") // 请求头部
		if origin != "" {
			c.ClientIP()
			// 接收客户端发送的origin （重要!）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			// 这是允许访问所有域 为 * 时
			// c.Header("Access-Control-Allow-Origin", "*")
			// 服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			// 允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, cache-Control, Content-Type, Pragma")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			// 设置缓存时间 有效期内不用再次发起 options 请求
			c.Header("Access-Control-Max-Age", "172800")
			// 允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		// OPTIONS方法 单独处理
		// 如果服务端支持该跨域请求，建议返回 204 状态码（返回 200 也可以）。204 不会返回任何内容可以节省一点数据传输
		// 如果不支持，建议返回 403 状态码（返回 404 或其他错误状态码也可以）
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
