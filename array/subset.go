package array

import "fmt"

func IsSubset(set []int, subset []int) bool {
	if len(set) < len(subset) || set[0] > subset[len(subset)-1] || set[len(set)-1] < subset[0] {
		return false
	}
	indset := make([]int, set[len(set)-1]+1)
	for i := 0; i < len(set); i++ {
		indset[set[i]]++
	}
	for i := 0; i < len(subset); i++ {
		if subset[i] > len(indset) {
			return false
		}
		indset[subset[i]]--
		if indset[subset[i]] < 0 {
			return false
		}
	}
	return true
}

func TestRun() {
	set := []int{2, 4, 4, 5, 6, 7, 8}
	subset1 := []int{9, 10, 11}
	subset2 := []int{1, 2}
	subset3 := []int{2, 5, 5, 6, 7}
	subset4 := []int{2, 3, 4, 5}

	fmt.Println("want false:", IsSubset(set, subset1))
	fmt.Println("want false:", IsSubset(set, subset2))
	fmt.Println("want true:", IsSubset(set, subset3))
	fmt.Println("want false:", IsSubset(set, subset4))
}
