package filter

import (
	"errors"
)

type SumFilter struct {
}

var SumFilterWongFormatError = errors.New("input data should be int array")

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]int)
	if !ok {
		return nil, ToIntFilterWongFormatError
	}
	var ret int
	for _, elem := range parts {
		ret += elem
	}
	return ret, nil
}
