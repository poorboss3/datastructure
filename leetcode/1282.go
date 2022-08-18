package leetcode

func groupThePeople(groupSizes []int) [][]int {
	res := make([][]int, 0)
	m := make(map[int][]int)
	for idx, size := range groupSizes {
		m[size] = append(m[size], idx)
	}
	for k, v := range m {
		for offset := 0; offset < len(v)/k; offset++ {
			res = append(res, v[offset*k:(offset+1)*k])
		}
	}
	return res
}
