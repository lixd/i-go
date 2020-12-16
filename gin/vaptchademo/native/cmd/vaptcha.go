package main

import (
	"i-go/gin/vaptchademo/native/logic"
	"net/http"
	"os"
	"path/filepath"
)

// http://localhost:8080/click.html
func main() {
	http.HandleFunc("/vaptcha/login", logic.Login)
	http.HandleFunc("/vaptcha/offline", logic.Offline)
	http.Handle("/", http.FileServer(http.Dir(buildPath())))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// buildPath build path to resource
func buildPath() (path string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir, _ = filepath.Split(dir)
	path = filepath.Join(dir, "static")
	return path
}
