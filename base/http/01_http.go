package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello,World \n")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":50051", nil)
	if err != nil {
		log.Fatal(err)
	}
}
