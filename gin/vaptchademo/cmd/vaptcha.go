package main

import (
	"github.com/lixd/vaptcha-sdk-go/examples/logic"
	"net/http"
)

func main() {
	http.HandleFunc("/vaptcha/login", logic.Login)
	http.HandleFunc("/vaptcha/offline", logic.Offline)
	http.HandleFunc("/click", logic.Click)
	http.HandleFunc("/invisible", logic.Invisible)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
