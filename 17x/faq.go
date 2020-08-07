package main

import (
	"fmt"
	"strings"
)

func main() {
	var CurrentJsVersion = "vaptcha-sdk-embed.2.8.7.js"
	CurrentJsVersion = strings.Replace(CurrentJsVersion, ".js", "-as.js", -1)
	fmt.Println(CurrentJsVersion)
}
