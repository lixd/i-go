package jwt

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const (
	hmacSampleSecret = "17x"
)

// GenerateToken 根据提供的信息生成 jwt token
func GenerateToken(userId int64) (string, error) {
	// 指定需要存储的数据
	claims := jwt.MapClaims{
		"userId": userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 秘钥签名
	return token.SignedString([]byte(hmacSampleSecret))
}

// ParseToken 从 jwt token 中解析出原始信息
func ParseToken(tokenString string) (userId int64, err error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	// 验证token,如果token被修改过则为false
	if !token.Valid {
		err = errors.New("invalid jwt token")
		return
	}
	if value, ok := claim["userId"]; ok {
		// 默认会处理成 float64
		userId = int64(value.(float64))
	}
	return
}

// secret 获取 secret 的方法 类型必须为 jwt.Keyfunc
func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// TODO 可以放在配置文件里
		return []byte(hmacSampleSecret), nil
	}
}
