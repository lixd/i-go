package utils

import (
	"github.com/satori/go.uuid"
)

type stringHelper struct {
}

var StringHelper = &stringHelper{}

/*
两个uuid库
https://github.com/google/uuid
https://github.com/satori/go.uuid
*/
func (stringHelper) GetUUID() (uuidHex string) {
	return uuid.NewV4().String()
}
