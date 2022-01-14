package jwt

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := Generate(int64(10086))
	if err != nil {
		fmt.Printf("err: %v \n", err.Error())
		return
	}
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDg2LCJhdWQiOiJ3ZWIiLCJleHAiOjE2NDE4ODYyODEsImp0aSI6ImRmM2ExOGM4N2Y1NDRlOWU4YzY4NDdlMGU4NGJjOWM5IiwiaWF0IjoxNjQxNzk5ODgxLCJpc3MiOiJpLWdvIiwibmJmIjoxNjQxNzk5ODgxLCJzdWIiOiJkZXZpY2UifQ.pfQIE3fKfayJJ1Xqf5ELvTmz_AIqZcD79SlbU9qtwqs"
	userId, err := Parse(tokenString)
	if err != nil {
		fmt.Printf("err: %v \n", err.Error())
		return
	}
	fmt.Println("userId: ", userId)
}

func BenchmarkGenerateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Generate(1)
	}
}

func BenchmarkParseToken(b *testing.B) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEyMzQ1fQ.vpZ2CycFtkbE63lSLyjG3y8mwtASRInPjynTh2be4Ks"
	for i := 0; i < b.N; i++ {
		_, _ = Parse(tokenString)
	}
}
