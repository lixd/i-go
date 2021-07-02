package main

import "fmt"

// options 模式 go-micro 使用的这种模式
// 对外提供 option 方法用于修改默认配置
// 而不是直接传动态参数之类的来修改
// 这样会更加灵活 就算参数有增减也不需要改动以前的代码
// 这种方式可以方便的提供默认值

// 另一种则是直接要求传一个 struct 为参数 etcd 使用的是这种模式
// struct 中包含了所有需要的参数
// 然后判断 如果某个值为空则赋为默认值
// option func for add custom config
type option func(*options)

type options struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Cli struct {
	opts   options
	others interface{}
}

// configure merge custom options
func configure(e *Cli, opts ...option) {
	for _, v := range opts {
		v(&e.opts)
	}
}

func main() {
	// custom config
	optionHost := func(i *options) {
		i.Host = "192.168.1.1"
	}
	cli := NewCli(optionHost)
	//	lasted config
	fmt.Println(cli)
}

// NewCli 需要的参数为 func，在 func 中执行对 配置的修改。
func NewCli(opts ...option) *Cli {
	// default config
	o := options{
		Host:     "127.0.0.1",
		Port:     "8080",
		Username: "guest",
		Password: "guest",
	}
	// new cli by default config
	cli := &Cli{
		opts:   o,
		others: nil,
	}
	//	merge config
	configure(cli, opts...)
	return cli
}
