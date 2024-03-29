package binarysearch

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
