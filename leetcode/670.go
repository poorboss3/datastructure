package leetcode

import "strconv"

func maximumSwap(num int) int {
	str := []byte(strconv.Itoa(num))
	start := 0
	maxIdx := 0
	for ; start < len(str)-1; start++ {
		max := int(str[start])
		maxIdx = start
		for i := start; i < len(str); i++ {
			if int(str[i]) >= max && int(str[i]) > int(str[start]) {
				max = int(str[i])
				maxIdx = i
			}
		}
		if maxIdx != start {
			break
		}
	}
	if start < maxIdx {
		str[start], str[maxIdx] = str[maxIdx], str[start]
	}
	ans, _ := strconv.ParseInt(string(str), 10, 32)
	return int(ans)
}
