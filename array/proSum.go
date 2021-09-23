package array

import (
	"fmt"
)

func ProSum(array []int) []int {
	proSum := make([]int, len(array)+1)
	proSum[0] = 0
	for i := 0; i < len(array); i++ {
		proSum[i+1] = proSum[i] + array[i]
	}
	return proSum
}

// SubSetSum 寻找数组里子数组和为sum的数量
func SubSetSum(array []int, sum int) int {
	var ans, sum0_i int = 0, 0
	proSum := make(map[int]int)
	proSum[0] = 1
	for i := 0; i < len(array); i++ {
		sum0_i += array[i]
		sum0_j := sum0_i - sum
		if _, ok := proSum[sum0_j]; ok {
			ans += proSum[sum0_j]
		}
		proSum[sum0_i]++
	}
	return ans
}

func Test_SubSetSum() {
	array := []int{3, 5, 2, -2, 4, 1}
	fmt.Println(SubSetSum(array, 0))
	fmt.Println(SubSetSum(array, 8))
	fmt.Println(SubSetSum(array, 5))
}

func DiffSet(array []int) []int {
	diff := make([]int, len(array))
	diff[0] = array[0]
	for i := 1; i < len(array); i++ {
		diff[i] = array[i] - array[i-1]
	}
	return diff
}

func ReDiffSet(diff []int) []int {
	array := make([]int, len(diff))
	array[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		array[i] = array[i-1] + diff[i]
	}
	return array
}

func Test_diffSet() {
	array := []int{2, 3, 4, 1, 5, 3, 5}
	fmt.Println(array)
	diff := DiffSet(array)
	fmt.Println(diff)
	diff[2] += 3
	diff[5] -= 3
	fmt.Println(ReDiffSet(diff))
}
