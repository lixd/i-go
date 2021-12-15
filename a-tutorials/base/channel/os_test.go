package channel

import (
	"fmt"
	"net/url"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("before")
	m.Run()
	fmt.Println("after")
}

func TestA(t *testing.T) {
	// rawUrl := "https://www.douban.com/accounts/connect/wechat/callback"
	rawUrl := "https://www.puug.com?url=https://www.puug.com"
	escape := url.QueryEscape(rawUrl)
	fmt.Println(escape)
}

func TestB(t *testing.T) {
	fmt.Println("B")
}
