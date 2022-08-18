package leetcode

import "strconv"

func solveEquation(equation string) string {
	equation = equation + "+"
	accountX, sum := 0, 0
	isLeft := 1
	add := 1
	nextNum := make([]byte, 0)
	for _, item := range []byte(equation) {
		switch item {
		case '=':
			num, _ := strconv.ParseInt(string(nextNum), 10, 64)
			if add*isLeft > 0 {
				sum += int(num)
			} else {
				sum -= int(num)
			}
			add = 1
			isLeft = -1
			nextNum = make([]byte, 0)
		case '+':
			num, _ := strconv.ParseInt(string(nextNum), 10, 64)
			if add*isLeft > 0 {
				sum += int(num)
			} else {
				sum -= int(num)
			}
			add = 1
			nextNum = make([]byte, 0)
		case '-':
			num, _ := strconv.ParseInt(string(nextNum), 10, 64)
			if add*isLeft > 0 {
				sum += int(num)
			} else {
				sum -= int(num)
			}
			add = -1
			nextNum = make([]byte, 0)
		case 'x':
			num := getNum(nextNum)
			if num == 0 {
			} else {
				if isLeft*add > 0 {
					accountX += int(num)
				} else {
					accountX -= int(num)
				}
			}
			nextNum = make([]byte, 0)

		default:
			nextNum = append(nextNum, item)
		}
	}
	if accountX == 0 && sum == 0 {
		return "Infinite solutions"
	}
	if accountX == 0 && sum != 0 {
		return "No solution"
	}
	return "x=" + strconv.Itoa(-sum/accountX)
}

func getNum(num []byte) int {
	if len(num) == 0 {
		return 1
	}
	n, _ := strconv.ParseInt(string(num), 10, 64)
	return int(n)
}
