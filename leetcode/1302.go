package leetcode

func deepestLeavesSum(root *TreeNode) int {
	nodeQueue := []*TreeNode{root}
	var sum int
	for len(nodeQueue) > 0 {
		sum = 0
		size := len(nodeQueue)
		for i := 0; i < size; i++ {
			node := nodeQueue[i]
			sum += nodeQueue[i].Val
			nodeQueue = nodeQueue[1:]
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
			}
		}
	}
	return sum
}
