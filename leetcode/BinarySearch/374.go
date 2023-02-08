package binarysearch

func guessNumber(n int) int {
	start := 0
	end := n
	for start <= end {
		target := (end-start)/2 + start
		gus := guess(target)
		if gus == 0 {
			return target
		} else if gus == -1 {
			end = target - 1
		} else {
			start = target + 1
		}
	}
	return -1
}

func guess(a int) int {
	return 0
}
