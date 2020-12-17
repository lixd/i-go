package tls

import (
	"os"
	"path"
)

const BasePath = "grpc/tls/assets"

func Server() (string, string) {
	dir, _ := os.Getwd()
	crt := path.Join(dir, BasePath+"/server.crt")
	key := path.Join(dir, BasePath+"/server.key")
	return crt, key
}
func Client() (string, string) {
	dir, _ := os.Getwd()
	crt := path.Join(dir, BasePath+"/client.crt")
	key := path.Join(dir, BasePath+"/client.key")
	return crt, key
}
