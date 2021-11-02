package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	r := "https://vaptcha.com/a/b/c/d/e/f"
	parse, err := url.Parse(r)
	if err != nil {
		return
	}
	parse2, err := url.Parse("./xxx.html")
	if err != nil {
		log.Println("err:", err)
	}
	reference := parse.ResolveReference(parse2)
	fmt.Println(reference)
}
