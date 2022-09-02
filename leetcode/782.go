package leetcode

func MovesToChessboard(board [][]int) int {
	changeCount := 0
	for ri := 0; ri < len(board)-1; ri++ {
		if isComplementary(board[ri], board[ri+1]) {
			continue
		} else {
			find := false
			for ni := ri + 2; ni < len(board); ni++ {
				if isComplementary(board[ri], board[ni]) {
					board[ri+1], board[ni] = board[ni], board[ri+1]
					changeCount++
					find = true
					break
				}
			}
			if !find {
				return -1
			}
		}
	}
	for ri := 0; ri < len(board)-1; ri++ {
		if isCiComplementary(board, ri, ri+1) {
			continue
		} else {
			findc := false
			for ci := ri + 2; ci < len(board); ci++ {
				if isCiComplementary(board, ri, ci) {
					board = changeColunms(board, ri+1, ci)
					changeCount++
					findc = true
					break
				}
			}
			if !findc {
				return -1
			}
		}
	}
	return changeCount
}

func changeColunms(board [][]int, col int, tarCol int) [][]int {
	for i := 0; i < len(board); i++ {
		board[i][col], board[i][tarCol] = board[i][tarCol], board[i][col]
	}
	return board
}

func isCiComplementary(board [][]int, col int, target int) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col]&board[i][target] == 0 {
			continue
		} else {
			return false
		}
	}
	return true
}

func isComplementary(col []int, nextCol []int) bool {
	for i := 0; i < len(col); i++ {
		if col[i]&nextCol[i] == 0 {
			continue
		} else {
			return false
		}
	}
	return true
}
