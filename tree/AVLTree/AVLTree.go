package avltree

type Item interface{}

type avlNode struct {
	value Item
	left  *avlNode
	right *avlNode
	high  int
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (n *avlNode) Height() int {
	if n == nil {
		return -1
	} else {
		return n.high
	}
}

func SingleRotateWithLeft(n *avlNode) *avlNode {
	var ln *avlNode
	ln = n.left
	n.left = ln.right
	ln.right = n
	n.high = max(n.left.Height(), n.right.Height()) + 1
	ln.high = max(ln.Height(), n.left.Height()) + 1
	return ln
}
