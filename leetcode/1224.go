package leetcode

func maxEqualFreq(nums []int) (ans int) {
	count := map[int]int{}
	freq := map[int]int{}
	maxFreq := 0
	for i, n := range nums {
		if count[n] > 0 {
			freq[count[n]]--
		}
		count[n]++
		maxFreq = max(maxFreq, count[n])
		freq[count[n]]++
		if maxFreq == 1 ||
			freq[maxFreq]*maxFreq+freq[maxFreq-1]*(maxFreq-1) == i+1 && freq[maxFreq] == 1 ||
			freq[maxFreq]*maxFreq+1 == i && freq[1] == 1 {
			ans = max(ans, i+1)
		}
	}
	return
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
