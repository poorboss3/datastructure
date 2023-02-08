package main

import (
	"fmt"
	"math"
	"strconv"
)

func getNum(num []byte) int {
	if len(num) == 0 {
		return 1
	}
	n, _ := strconv.ParseInt(string(num), 10, 64)
	return int(n)
}

func main() {
	x := math.Inf(1)
	switch {
	case x > 0, x < 0:
		fmt.Println(x)
	case x == 0:
		fmt.Println("0")
	default:
		fmt.Println("else")
	}
}

/**
["1","1","0","0","0"]
["1","1","0","0","0"]
["0","0","1","0","0"]
["0","0","0","1","1"]
**/

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

func openLock(deadends []string, target string) int {
	setp := 0
	for _, s := range target[:len(target)-1] {
		n, _ := strconv.ParseInt(string(s), 10, 64)
		if n == 5 {
			setp += 5
		} else {
			setp += int(n % 5)
		}
	}
	last, _ := strconv.ParseInt(string(target[len(target)-1:]), 10, 64)
	forward := last <= 5
	backword := last > 5
	if forward && !container(deadends, string(target[:len(target)-1]), 1, int(last)) {
		setp += int(last)
	} else {
		backword = true
	}
	if backword && !container(deadends, string(target[:len(target)-1]), int(last), 9) {
		setp += (9 - int(last))
	} else {
		return -1
	}
	return setp
}

func container(array []string, target string, start, end int) bool {
	for _, item := range array {
		for i := start; i <= end; i++ {
			if item == target+strconv.Itoa(i) {
				return true
			}
		}
	}
	return false
}

func trimMean(arr []int) float64 {
	fmt.Println(len(arr))
	sortArr := QuickSort(arr)
	sum := 0
	cap := int(math.Floor(float64(len(sortArr)) * 0.05))
	for i := cap; i < len(sortArr)-cap; i++ {
		sum += arr[i]
	}
	fmt.Println(sum, "/", len(arr)-cap)
	return float64(sum) / float64((len(arr) - 2*cap))
}

func QuickSort(array []int) []int {
	l := len(array)
	if l <= 1 {
		return array
	}
	temp := make([]int, 0)
	pivot := findPivot(array, 0, l-1, l/2)
	pivotValue := array[pivot]
	i := 0
	j := l - 2
	array[pivot] = array[l-1]
	array[l-1] = pivotValue
	for i <= j {
		if array[i] < array[l-1] {
			i++
		} else {
			array[i], array[j] = array[j], array[i]
			j--
		}
	}
	array[l-1] = array[i]
	array[i] = pivotValue
	temp = append(temp, QuickSort(array[:i])...)
	temp = append(temp, pivotValue)
	temp = append(temp, QuickSort(array[i+1:])...)
	return temp
}

func findPivot(array []int, left int, right int, center int) int {
	if (array[left]-array[center])*(array[center]-array[right]) > 0 {
		return center
	} else if (array[center]-array[left])*(array[left]-array[right]) > 0 {
		return left
	} else {
		return right
	}
}
