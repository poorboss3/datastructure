package avltree

import (
	"fmt"
	"sync"
)

type Item interface{}

type avlNode struct {
	value Item
	key   int
	left  *avlNode
	right *avlNode
	high  int
}

type AvlTree struct {
	root *avlNode
	lock sync.RWMutex
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

func singleRotateWithLeft(n *avlNode) *avlNode {
	var ln *avlNode
	ln = n.left
	n.left = ln.right
	ln.right = n
	n.high = max(n.left.Height(), n.right.Height()) + 1
	ln.high = max(ln.Height(), n.left.Height()) + 1
	return ln
}

func doubleRotateWithLeft(n *avlNode) *avlNode {
	n.left = singleRotateWithLeft(n.left)

	return singleRotateWithLeft(n)
}

func singleRotateWithRight(n *avlNode) *avlNode {
	var ln *avlNode
	ln = n.right
	n.right = ln.left
	ln.left = n
	n.high = max(n.left.Height(), n.right.Height()) + 1
	ln.high = max(ln.Height(), n.right.Height()) + 1
	return ln
}

func doubleRotateWithRight(n *avlNode) *avlNode {
	n.right = singleRotateWithRight(n.right)

	return singleRotateWithRight(n)
}

func insert(n, newNode *avlNode) {
	if n.key > newNode.key {
		insert(n.left, newNode)
		if n.left.Height()-n.right.Height() == 2 {
			if newNode.key < n.left.key {
				n = singleRotateWithLeft(n)
			} else {
				n = doubleRotateWithLeft(n)
			}
		}
	} else if n.key < newNode.key {
		insert(n.right, newNode)
		if n.right.Height()-n.left.Height() == 2 {
			if newNode.key > n.right.key {
				n = singleRotateWithRight(n)
			} else {
				n = doubleRotateWithRight(n)
			}
		}
	}
	n.high = max(n.left.Height(), n.right.Height()) + 1
}

func (avlt *AvlTree) Insert(value Item, key int) {
	node := &avlNode{value: value, key: key, left: nil, right: nil, high: 0}
	avlt.lock.Lock()
	defer avlt.lock.Unlock()
	if avlt.root == nil {
		avlt.root = node
		return
	}
	insert(avlt.root, node)
}

func (avlt *AvlTree) String() {
	avlt.lock.Lock()
	defer avlt.lock.Unlock()
	fmt.Println("--------------------------")
	stringify(avlt.root, 0)
	fmt.Println("--------------------------")
}

func stringify(n *avlNode, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "    "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}
