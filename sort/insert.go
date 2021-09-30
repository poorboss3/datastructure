package sort

func InsertSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	var j int
	for i := 1; i < len(array); i++ {
		temp := array[i]
		for j = i; j > 0 && array[j-1] > temp; j-- {
			array[j] = array[j-1]
		}
		array[j] = temp
	}
	return array
}
