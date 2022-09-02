package leetcode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	biggest := 0
	for i, n := range nums {
		if n > nums[biggest] {
			biggest = i
		}
	}
	return &TreeNode{
		Val:   nums[biggest],
		Left:  constructMaximumBinaryTree(nums[:biggest]),
		Right: constructMaximumBinaryTree(nums[biggest+1:]),
	}
}
