package binarysearch

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	left, riht := 0, len(arr)-1
	for left <= riht {
		mid := (riht-left)/2 + left
		fmt.Println(mid)
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] > arr[mid-1] {
			left = mid + 1
		} else {
			riht = mid - 1
		}
	}
	return -1
}
