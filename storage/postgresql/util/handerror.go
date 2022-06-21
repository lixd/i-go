package util

import (
	"log"
)

func HandError(msg string, err error) {
	if err != nil {
		log.Printf("msg:%s err:%v", msg, err)
	}
}
