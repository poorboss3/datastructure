package leetcode

func findLongestChain(pairs [][]int) int {
	pairs = QuickSort(pairs)
	cur := ^int(^uint(0) >> 1)
	ans := 0
	for _, item := range pairs {
		if item[0] > cur {
			cur = item[1]
			ans++
		}
	}
	return ans
}

func QuickSort(pairs [][]int) [][]int {
	if len(pairs) <= 1 {
		return pairs
	}
	piovt := findPiovt(pairs, 0, len(pairs)-1, len(pairs)/2)
	pivotValue := pairs[piovt]
	pairs[len(pairs)-1], pairs[piovt] = pairs[piovt], pairs[len(pairs)-1]
	i := len(pairs) - 2
	j := 0
	for i >= j {
		if pairs[j][1] < pairs[len(pairs)-1][1] {
			j++
		} else {
			pairs[j], pairs[i] = pairs[i], pairs[j]
			i--
		}
	}
	pairs[len(pairs)-1], pairs[j] = pairs[j], pairs[len(pairs)-1]
	temp := make([][]int, 0)
	temp = append(temp, QuickSort(pairs[:j])...)
	temp = append(temp, pivotValue)
	temp = append(temp, QuickSort(pairs[j+1:])...)
	return temp
}

func findPiovt(pairs [][]int, left, right, center int) int {
	if (pairs[left][1]-pairs[center][1])*(pairs[center][1]-pairs[right][1]) > 0 {
		return center
	} else if (pairs[center][1]-pairs[left][1])*(pairs[left][1]-pairs[right][1]) > 0 {
		return left
	} else {
		return right
	}
}
