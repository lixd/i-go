package filter

import (
	"errors"
	"strings"
)

type SplitFilter struct {
	delimiter string
}

var SplitFilterWongFormatError = errors.New("input data should be string")

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter: delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, SplitFilterWongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
