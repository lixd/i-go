package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type hello struct {
}

var Hello = hello{}

// Greeter godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /accounts/{id} [get]
func (hello) Greeter(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		name = "world"
	}
	c.JSON(http.StatusOK, "hello"+name)
}
