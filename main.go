package main

//array "arithmeticand/array"

import "arithmeticand/weichatdat"

var (
	a = b + c
	b = f("b")
	c = f("c")
	d = 3
)

func f(caller string) int {
	d++
	return d
}

func main() {
	// heap := heap.NewLeftHeap(10)
	// for _, item := range []int{5, 3, 7, 9, 2, 1, 6, 4, 8} {
	// 	heap = heap.Insert(item)
	// }

	// for i := 0; i < 9; i++ {
	// 	min := heap.DeleteMin()
	// 	fmt.Printf("%d \n", min)
	// }
	weichatdat.DatToImage("E:\\wechat\\Documents\\WeChat Files\\poorboss123\\FileStorage\\Image\\", "F:\\1111")
}
