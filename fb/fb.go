package fb

//FbLoop is the constructor of Cache
func FbLoop(len int) int {
	lastVal := 1
	llastVal := 1
	curVal := 1
	for i := 2; i < len; i++ {
		curVal = lastVal + llastVal
		llastVal = lastVal
		lastVal = curVal
	}
	return curVal
}

//Fbrecursion is the constructor of Cache
func Fbrecursion(len int) int {
	if len <= 2 {
		return 1
	}
	return Fbrecursion(len-2) + Fbrecursion(len-1)
}
