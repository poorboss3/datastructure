package heap

import "fmt"

type LeftHeap struct {
	Data  int
	Left  *LeftHeap
	Right *LeftHeap
	Npl   int
}

func NewLeftHeap(root int) *LeftHeap {
	lh := &LeftHeap{Data: root, Left: nil, Right: nil, Npl: 0}
	return lh
}

func (lh *LeftHeap) Insert(value int) *LeftHeap {
	newNode := &LeftHeap{Data: value, Left: nil, Right: nil, Npl: 0}
	return Merge(lh, newNode)
}

func (lh *LeftHeap) DeleteMin() int {
	min := lh.Data
	left := lh.Left
	right := lh.Right
	m := Merge(left, right)
	lh.Data = m.Data
	lh.Left = m.Left
	lh.Right = m.Right
	lh.Npl = m.Npl
	m = nil
	return min
}

func Merge(h1 *LeftHeap, h2 *LeftHeap) *LeftHeap {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	var m *LeftHeap
	if h1.Data < h2.Data {
		m = merge1(h1, h2)
	} else {
		m = merge1(h2, h1)
	}
	return m
}

func merge1(h1 *LeftHeap, h2 *LeftHeap) *LeftHeap {
	if h1.Left == nil {
		h1.Left = h2
	} else {
		h1.Right = Merge(h1.Right, h2)
		if h1.Left.Npl < h1.Right.Npl {
			temp := h1.Right
			h1.Right = h1.Left
			h1.Left = temp
		}
		h1.Npl = h1.Right.Npl + 1
	}
	return h1
}

func (lh *LeftHeap) Print() {
	if lh != nil {
		fmt.Printf("%d ", lh.Data)
		lh.Left.Print()
		lh.Right.Print()
	}
}
