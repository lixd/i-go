package main

import (
	"github.com/gin-gonic/gin"
)

var swagHandler gin.HandlerFunc

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server.
// @termsOfService  https://lixueduan.com

// @contact.name   lixd
// @contact.url    https://lixueduan.com
// @contact.email  xueduan.li@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// SwaggerUI: http://localhost:8080/swagger/index.html
func main() {

	e := gin.Default()

	v1 := e.Group("/api/v1")
	{
		v1.GET("/hello", Hello)
		v1.POST("/login", Login)
	}

	if swagHandler != nil {
		e.GET("/swagger/*any", swagHandler)
	}

	if err := e.Run(":8080"); err != nil {
		panic(err)
	}
}
