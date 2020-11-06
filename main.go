package main

import (
	tree "arithmeticand/tree"
	"fmt"
)

func main() {
	TreeNode := tree.TreeNode{12, nil, nil}
	//var t tree.Tree
	t := &TreeNode
	t.Value = 14
	fmt.Printf("oldvalue is %d,newvalue is %d", TreeNode.Value, t.Value)
}
 