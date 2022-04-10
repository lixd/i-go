package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// 生成证书 go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	message := fmt.Sprintf("Hello %s!", r.UserAgent())
	w.Write([]byte(message))
}

func main() {
	h2()
	// server := &http.Server{
	// 	Addr: ":9876",
	// 	Handler: h2c.NewHandler(
	// 		http.HandlerFunc(indexHandler),
	// 		&http2.Server{},
	// 	),
	// }
	//
	// if err := server.ListenAndServe(); err != nil {
	// 	panic(err)
	// }
}

func h2() {
	server := &http.Server{
		Addr: ":9875",
		Handler: h2c.NewHandler(
			http.HandlerFunc(indexHandler),
			&http2.Server{},
		),
	}

	if err := server.ListenAndServeTLS("assets/cert.pem", "assets/key.pem"); err != nil {
		panic(err)
	}
}
