package main

import (
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func HelloServerPprof(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello,World \n")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/hello", HelloServerPprof)
	err := http.ListenAndServe(":50051", nil)
	if err != nil {
		log.Fatal(err)
	}
}
