package array

//MaxSubsequenceSum is method acc the max subsequen
func MaxSubsequenceSum(a []int, N int) int {
	var ThisSum, MaxSum int
	MaxSum = 0
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			ThisSum = 0
			for k := i; k <= j; k++ {
				ThisSum += a[k]
			}
			if ThisSum > MaxSum {
				MaxSum = ThisSum
			}
		}
	}
	return MaxSum
}

//MaxSubsequenceSumDi is balabala...
func MaxSubsequenceSumDi(a []int, N int) int {
	var thisSum, maxSum int
	maxSum = 0
	for i := 0; i < N; i++ {
		thisSum = 0
		for j := i; j < N; j++ {
			thisSum += a[j]
			if thisSum > maxSum {
				maxSum = thisSum
			}
		}
	}
	return maxSum
}

//MaxSubsequenceSumDia is balabala...
func MaxSubsequenceSumDia(a []int, N int) int {
	return maxSub(a, 0, N)
}

func maxSub(a []int, left int, right int) int {
	var maxLeftSum, maxRightSum int
	var maxLeftBorderSum, maxRightBorderSum int
	var leftBorderSum, rightBorderSum int
	var center int

	if left == right {
		if a[left] > 0 {
			return a[left]
		}
		return 0
	}

	center = (left + right) / 2
	maxLeftSum = maxSub(a, left, center)
	maxRightSum = maxSub(a, center+1, right)
	maxLeftBorderSum = 0
	leftBorderSum = 0
	for i := center; i >= left; i-- {
		leftBorderSum += a[i]
		if leftBorderSum > maxLeftBorderSum {
			maxLeftBorderSum = leftBorderSum
		}
	}
	maxRightBorderSum = 0
	rightBorderSum = 0
	for i := center + 1; i <= right; i++ {
		rightBorderSum += a[i]
		if rightBorderSum > maxRightBorderSum {
			maxRightBorderSum = rightBorderSum
		}
	}
	if maxLeftBorderSum+maxRightBorderSum > maxLeftSum && maxLeftBorderSum+maxRightBorderSum > maxRightSum {
		return maxLeftBorderSum + maxRightBorderSum
	} else if maxLeftSum > maxRightSum {
		return maxLeftSum
	} else {
		return maxRightSum
	}
}

//MaxSubsequenceSumOn is balabala
func MaxSubsequenceSumOn(a []int, N int) int {
	//var thisSum, maxSum int
	thisSum := 0
	maxSum := 0
	for i := 0; i < N; i++ {
		thisSum += a[i]
		if thisSum > maxSum {
			maxSum = thisSum
		} else if thisSum < 0 {
			thisSum = 0
		}
	}
	return maxSum
}
