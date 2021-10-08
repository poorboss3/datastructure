package sort

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	center := len(array) / 2
	left := MergeSort(array[:center])
	right := MergeSort(array[center:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	lLen := len(left)
	rLen := len(right)
	merge := make([]int, 0)
	lptr, rptr := 0, 0
	for lptr < lLen && rptr < rLen {
		if left[lptr] <= right[rptr] {
			merge = append(merge, left[lptr])
			lptr++
		} else {
			merge = append(merge, right[rptr])
			rptr++
		}
	}
	if lptr < lLen {
		merge = append(merge, left[lptr:]...)
	}
	if rptr < rLen {
		merge = append(merge, right[rptr:]...)
	}
	return merge
}
