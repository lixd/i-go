package jwt

import (
	"time"

	"i-go/utils"

	"github.com/golang-jwt/jwt"
)

// https://jwt.io/ 在线解码
// https://datatracker.ietf.org/doc/html/rfc7519 rfc 文档

/*
JWT 优点是无状态的
确定则是一旦发布后无法撤销，比如用户改密码后旧的JWT任然可以使用
一般使用JWT黑名单方式处理，比如退出登录或者改密码后把之前的JWT写入黑名单
*/

const (
	ExpireTime = time.Hour * 24 // JWT 有效期24小时
)

// CustomClaims 自定义Claims
type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

// Generate 生成jwt
func Generate(userId int64) (string, error) {
	claim := CustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Audience:  "web",                             // 观众，相当于接受者
			ExpiresAt: time.Now().Add(ExpireTime).Unix(), // 过期时间
			Id:        utils.StringHelper.GetUUID(),      // jwt 编号
			IssuedAt:  time.Now().Unix(),                 // 发布时间
			Issuer:    "i-go",                            // 发布者
			NotBefore: time.Now().Unix(),                 // 生效时间，在生效时间前该JWT依旧无效
			Subject:   "device",                          // 主题
		},
	}
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return st.SignedString(secret())
}

func keyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return secret(), nil
	}
}

func secret() []byte {
	return []byte("my secret")
}

// Parse 解析jwt
func Parse(token string) (CustomClaims, error) {
	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, keyFunc())
	if err != nil {
		return CustomClaims{}, err
	}
	// 类型转换并判断是否有效
	if claim, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return *claim, nil
	}
	return CustomClaims{}, err
}
