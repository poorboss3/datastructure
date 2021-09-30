package heap

import (
	"errors"
	"fmt"
)

type MinHeap struct {
	Len      int
	Capacity int
	Data     []int
}

func NewHeap(capacity int) (*MinHeap, error) {
	if capacity < 0 {
		return nil, errors.New("capacity must bigger then zore")
	}
	heap := &MinHeap{Capacity: capacity, Len: 0}
	heap.Data = make([]int, capacity+1)
	return heap, nil
}

func (h *MinHeap) Insert(value int) error {
	var i int
	if h.Len == h.Capacity {
		return errors.New("heap was full")
	}
	if h.Len == 0 {
		h.Data[1] = value
		h.Len++
		return nil
	}
	for i = h.Len + 1; i > 0 && h.Data[i/2] > value; i /= 2 {
		h.Data[i] = h.Data[i/2]
	}
	h.Data[i] = value
	h.Len++
	return nil
}

func (h *MinHeap) DeleteMin() (int, error) {
	if h.Len == 0 {
		return 0, errors.New("heap was null")
	}
	min := h.Data[1]
	child := 1
	for i := 1; i <= h.Len/2; i = child {
		if h.Data[2*i] > h.Data[2*i+1] && h.Data[2*i+1] > 0 {
			h.Data[i] = h.Data[2*i+1]
			child = 2*i + 1
		} else {
			h.Data[i] = h.Data[2*i]
			child = 2 * i
		}
	}
	h.Data[child] = h.Data[h.Len]
	h.Data[h.Len] = 0
	h.Len--
	return min, nil
}

func (h *MinHeap) Print() {
	fmt.Printf("len:%d,cap:%d,data:%v\n", h.Len, h.Capacity, h.Data)
}
