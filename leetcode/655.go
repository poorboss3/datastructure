package leetcode

import "strconv"

// Definition for a binary tree node.
func printTree(root *TreeNode) [][]string {
	height := deepLen(root)
	m := height + 1
	n := 1<<m - 1
	ans := make([][]string, m)
	for i := range ans {
		ans[i] = make([]string, n)
	}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, r, c int) {
		ans[r][c] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfs(node.Left, r+1, c-1<<(height-r-1))
		}
		if node.Right != nil {
			dfs(node.Right, r+1, c+1<<(height-r-1))
		}
	}
	dfs(root, 0, (n-1)/2)
	return ans
}

func deepLen(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil {
		return deepLen(root.Left) + 1
	} else {
		return max(deepLen(root.Left)+1, deepLen(root.Right)+1)
	}
}
