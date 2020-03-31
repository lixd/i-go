package simple

import (
	"fmt"
	"net/http"
)

func main() {
	// 1.存到路由表中 map结构 pattern为key value就是func
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	// 内部调用net.Listen("tcp", addr)初始化socket, bind, listen的操作.
	// 同时根据请求的host+url匹配路由表中的func 没有则返回 NotFoundHandler 404
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("start http server fail:", err)
	}
}
