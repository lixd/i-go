package pattern

/*
通过这种options模式，可以不必每次定义所有的选项，只需选择自己想要的改动即可。
*/

type Server struct {
	cfg config
}

// 配置项
type config struct {
	Tracing bool
}

// 配置Option的包装函数
type Option func(*Server)

// 添加开启tracing的可选项
func WithTracing() Option {
	return func(r *Server) {
		r.cfg.Tracing = true
	}
}

// 使用可选项进行构建
func New(name string, opts ...Option) *Server {
	const (
		defaultTracing = false
	)

	server := &Server{
		cfg: config{
			Tracing: defaultTracing, // 进行默认的初始化赋值
		},
	}

	// 查看是否有可选项，如果有则使用可选项将默认值覆盖。
	for _, opt := range opts {
		opt(server)
	}
	return server
}
