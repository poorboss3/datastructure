package main

import (
	avltree "arithmeticand/tree/avltree"
	btree "arithmeticand/tree/binarysearchtree"
)

var bst btree.ItemBinarySearchTree
var avlt avltree.AvlTree

func fillTree(bst *btree.ItemBinarySearchTree) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
}

func fillAvlTree(avlt *avltree.AvlTree) {
	for i := 1; i < 20; i++ {
		avlt.Insert(i, i)
	}
}

func main() {
	fillAvlTree(&avlt)
	avlt.String()
}
