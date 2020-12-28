package main

import (
	"log"
	"net/http"

	"i-go/gin/vaptchademo/native/logic"
)

// http://localhost:8080/vaptcha/demo/click.html
func main() {
	http.HandleFunc("/vaptcha/login", logic.Login)
	http.HandleFunc("/vaptcha/offline", logic.Offline)
	http.Handle("/vaptcha/demo/", http.StripPrefix("/vaptcha/demo", http.FileServer(http.Dir("../../assets"))))
	log.Println("server run on 0.0.0.0:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Run error:%v", err)
	}
}
