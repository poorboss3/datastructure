package leetcode

//
//     for cur := root; cur != nil; cur = cur.Right {
//         if val > cur.Val {
//             if parent == nil {
//                 return &TreeNode{val, root, nil}
//             }
//             parent.Right = &TreeNode{val, cur, nil}
//             return root
//         }
//         parent = cur
//     }
//     parent.Right = &TreeNode{Val: val}
//     return root
// }
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	for cur := root; cur != nil; cur = cur.Right {
		if val > cur.Val {
			if parent == nil {
				return &TreeNode{Val: val, Left: root}
			}
			parent.Right = &TreeNode{val, cur, nil}
			return root
		}
		parent = cur
	}
	parent.Right = &TreeNode{Val: val}
	return root
}
