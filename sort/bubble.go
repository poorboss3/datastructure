package sort

func BubbleSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	for i := 0; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				array[j] = array[j] + array[j-1]
				array[j-1] = array[j] - array[j-1]
				array[j] = array[j] - array[j-1]
			}
		}
	}
	return array
}
