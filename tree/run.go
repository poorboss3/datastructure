package main

import (
	btree "arithmeticand/tree/binarysearchtree"
	"fmt"
)

var bst btree.ItemBinarySearchTree

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

func main() {
	fillTree(&bst)
	var pre, in, post []string
	bst.PreorderTraverse(func(i btree.Item) {
		pre = append(pre, fmt.Sprintf("%s", i))
	})
	bst.InorderTraverse(func(i btree.Item) {
		in = append(in, fmt.Sprintf("%s", i))
	})
	bst.PostorderTraverse(func(i btree.Item) {
		post = append(post, fmt.Sprintf("%s", i))
	})
	fmt.Printf("pre order is:%v\n", pre)
	fmt.Printf("in order is:%v\n", in)
	fmt.Printf("post order is:%v\n", post)
}
