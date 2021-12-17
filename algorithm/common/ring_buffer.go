package common

import "github.com/pkg/errors"

type ringBuffer struct {
	data    []int64
	r, w, l int64
}

func NewRingBuffer(l int64) *ringBuffer {
	return &ringBuffer{
		data: make([]int64, l),
		r:    0,
		w:    0,
		l:    l,
	}
}

func (r *ringBuffer) Read() (int64, error) {
	if r.r == r.w {
		return 0, errors.New("buffer is empty")
	}
	data := r.data[r.r%r.l]
	r.r++
	return data, nil
}

func (r *ringBuffer) Write(data int64) error {
	if r.w-r.r == r.l {
		return errors.New("buffer is full")
	}
	r.data[r.w%r.l] = data
	r.w++
	return nil
}
