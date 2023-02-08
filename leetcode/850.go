package leetcode

import (
	"fmt"
	"math"
)

func rectangleArea(rectangles [][]int) int {
	xmap := make(map[int][]int)
	for _, rectange := range rectangles {
		for i := rectange[0]; i <= rectange[2]; i++ {
			if xval, ok := xmap[i]; ok {
				if xval[0] > rectange[1] {
					xval[0] = rectange[1]
				}
				if xval[1] < rectange[3] {
					xval[1] = rectange[3]
				}
			} else {
				xmap[i] = make([]int, 2)
				xmap[i][0] = rectange[1]
				xmap[i][1] = rectange[3]
			}
		}
	}
	fmt.Println(xmap)
	sum := 0
	for _, value := range xmap {
		sum += value[1] - value[0]
	}
	return int(math.Mod(float64(sum), math.Pow10(9)+7))
}
