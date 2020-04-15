package utils

import (
	uuid "github.com/satori/go.uuid"
)

type stringHelper struct {
}

var StringHelper = &stringHelper{}

func (stringHelper) GetUUID() (uuidHex string) {
	return uuid.NewV4().String()
}
