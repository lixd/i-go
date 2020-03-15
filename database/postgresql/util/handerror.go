package util

import (
	"github.com/prometheus/common/log"
)

func HandError(msg string, err error) {
	if err != nil {
		log.Error(msg, err)
	}
}
