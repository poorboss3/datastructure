package binarysearchtree

import (
	"fmt"
	"sync"
)

type Item interface{}

type Node struct {
	key   int
	value Item
	left  *Node
	right *Node
}

type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (bst *ItemBinarySearchTree) Insert(key int, value Item) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node{key: key, value: value, left: nil, right: nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

func (bst *ItemBinarySearchTree) Min() *Node {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	return findmin(bst.root)
}

func findmin(node *Node) *Node {
	if node == nil {
		return nil
	}
	for {
		if node.left == nil {
			return node
		}
		node = node.left
	}
}

func (bst *ItemBinarySearchTree) Max() *Node {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	return findmax(bst.root)
}
func findmax(node *Node) *Node {
	if node == nil {
		return nil
	}
	for {
		if node.right == nil {
			return node
		}
		node = node.right
	}
}

func (bst *ItemBinarySearchTree) Search(key int) *Node {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	return searchNode(bst.root, key)
}

func searchNode(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	if n.key > key {
		return searchNode(n.left, key)
	} else if n.key < key {
		return searchNode(n.right, key)
	} else {
		return n
	}
}

// internal recursive function to remove an item
func delete(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key < key {
		node.right = delete(node.right, key)
		return node
	} else if node.key > key {
		node.left = delete(node.left, key)
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	} else if node.left == nil {
		node = node.right
		return node
	} else if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	minNode := node.right
	for {
		if minNode != nil && minNode.left != nil {
			minNode = minNode.left
		} else {
			break
		}
	}
	node.key, node.value = minNode.key, minNode.value
	node.right = delete(node.right, node.key)
	return node
}

func (bst *ItemBinarySearchTree) Delete(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	delete(bst.root, key)
}

func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("--------------------------")
	stringify(bst.root, 0)
	fmt.Println("--------------------------")
}

func stringify(n *Node, level int) {
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

func (bst *ItemBinarySearchTree) InorderTraverse(f func(Item)) {

	inorderTraverse(bst.root, f)
}

func inorderTraverse(n *Node, f func(Item)) {
	if n != nil {
		inorderTraverse(n.left, f)
		f(n.value)
		inorderTraverse(n.right, f)
	}
}

func (bst *ItemBinarySearchTree) PreorderTraverse(f func(Item)) {

	preorderTraverse(bst.root, f)
}

func preorderTraverse(n *Node, f func(Item)) {
	if n != nil {
		f(n.value)
		preorderTraverse(n.left, f)
		preorderTraverse(n.right, f)
	}
}

func (bst *ItemBinarySearchTree) PostorderTraverse(f func(Item)) {

	postorderTraverse(bst.root, f)
}

func postorderTraverse(n *Node, f func(Item)) {
	if n != nil {
		postorderTraverse(n.left, f)
		postorderTraverse(n.right, f)
		f(n.value)
	}
}
