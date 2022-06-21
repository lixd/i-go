//go:build doc

package main

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "i-go/gin/swagger/docs"
)

func init() {
	swagHandler = ginSwagger.WrapHandler(files.Handler)
}
