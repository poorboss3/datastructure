package leetcode

func findClosestElements(arr []int, k int, x int) []int {
	if arr[0] >= x {
		return arr[:k]
	}
	if arr[len(arr)-1] <= x {
		return arr[len(arr)-1-k:]
	}
	i := 0
	for ; i+k < len(arr); i++ {
		if abs(arr[i]-x) < abs(arr[i+k]-x) {
			break
		}
		if abs(arr[i]-x) == abs(arr[i+k]-x) && arr[i+k] > arr[i] {
			break
		}
	}
	return arr[i:(i + k)]
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
