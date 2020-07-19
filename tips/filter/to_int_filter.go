package filter

import (
	"errors"
	"strconv"
)

type ToIntFilter struct {
}

var ToIntFilterWongFormatError = errors.New("input data should be string array")

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tif *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWongFormatError
	}
	var ret []int
	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, i)
	}
	return ret, nil
}
