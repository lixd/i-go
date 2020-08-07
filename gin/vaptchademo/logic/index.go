package logic

import (
	"fmt"
	"net/http"
	"os"
)

func Click(writer http.ResponseWriter, request *http.Request) {
	wd, _ := os.Getwd()
	wd = wd + "/examples/static/click.html"
	fmt.Println(wd)
	http.ServeFile(writer, request, wd)
}
func Invisible(writer http.ResponseWriter, request *http.Request) {
	wd, _ := os.Getwd()
	wd = wd + "/examples/static/invisible.html"
	fmt.Println(wd)
	http.ServeFile(writer, request, wd)
}
