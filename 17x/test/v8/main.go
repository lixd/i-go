package main

import (
	"github.com/lazytiger/go-v8"
)

/*
需要安装v8环境 具体见:https://github.com/lazytiger/go-v8/blob/master/install.sh
*/
func main() {
	engine := v8.NewEngine()
	script := engine.Compile([]byte("'Hello ' + 'World!'"), nil, nil)
	context := engine.NewContext(nil)

	context.Scope(func(cs v8.ContextScope) {
		result := script.Run()
		println(result.ToString())
	})
}
