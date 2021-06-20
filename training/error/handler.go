package main

import (
	"encoding/json"
	"io"
	"log"
)

type Config struct {
}

// WriteAll 错误做法: 打印日志后 又再次将 error 返回，上一层检测到错误后可能也会打印日志并再往上层返回，最终导致一个 error 打印了 N 个日志。
func WriteAll(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if err != nil {
		log.Println("unable to write:", err) // annotated error
		return err                           // unannotated error
	}
	return nil
}
func WriteConfig(w io.Writer, conf *Config) error {
	buf, err := json.Marshal(conf)
	if err != nil {
		log.Printf("could not marshal config: %v", err)
		// return err
		// oops，forgot to return
	}
	if err := WriteAll(w, buf); err != nil {
		log.Println("could not write config: %v", err)
		return err
	}
	return nil
}
