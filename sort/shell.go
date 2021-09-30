package sort

import "fmt"

func ShellSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	var j int
	for inc := len(array) / 2; inc > 0; inc /= 2 {
		for i := 1; i*inc < len(array); i++ {
			temp := array[i*inc]
			for j = i; j > 0 && array[(j-1)*inc] > temp; j-- {
				array[j*inc] = array[(j-1)*inc]
			}
			array[j*inc] = temp
		}
		fmt.Println(array)
	}
	return array
}
