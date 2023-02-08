package leetcode

import (
	"container/list"
	"math"
)

type Index struct {
	X int
	Y int
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	items := list.New()
	if n%2 == 0 {
		items.PushBack(Index{n / 2, n / 2})
		items.PushBack(Index{n/2 - 1, n / 2})
		items.PushBack(Index{n / 2, n/2 - 1})
		items.PushBack(Index{n/2 - 1, n/2 - 1})
	} else {
		items.PushBack(Index{n / 2, n / 2})
	}
	for item := items.Front(); item != nil; item = item.Next() {
		idx := item.Value.(Index)
		if ok, plus := maxPlus(idx.X, idx.Y, mines, n); ok {
			return plus
		}
	}
	return 0
}

func maxPlus(x int, y int, mines [][]int, n int) (bool, int) {
	plus := min([]int{x, y, n - 1 - x, n - 1 - y})
	for _, mine := range mines {
		if (mine[0] == x && int(math.Abs(float64(y-mine[1]))) <= plus) || (mine[1] == y && int(math.Abs(float64(x-mine[0]))) <= plus) {
			return false, 0
		}
	}
	return true, plus + 1
}
