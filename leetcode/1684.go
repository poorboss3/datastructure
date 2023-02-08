package leetcode

func countConsistentStrings(allowed string, words []string) int {
	allowMap := make(map[rune]bool)
	for _, s := range allowed {
		allowMap[s] = true
	}
	res := 0
	for _, word := range words {
		exist := true
		for _, w := range word {
			if _, ok := allowMap[w]; !ok {
				exist = false
				break
			}
		}
		if exist {
			res++
		}
	}
	return res
}
