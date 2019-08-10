package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserRegister struct {
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
}
type V5Controller struct {
}

func (v5Controller *V5Controller) Register(c *gin.Context) {
	var user UserRegister
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, "参数错误")
		return
	}
	log.Println(user.Email, user.Password, user.PasswordAgain)
}
