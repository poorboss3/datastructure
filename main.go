package main

import (
	"arithmeticand/leetcode"
	"fmt"
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
	fmt.Println(leetcode.ValidateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(leetcode.ValidateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
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
