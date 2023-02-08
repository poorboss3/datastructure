package leetcode

func reorderSpaces(text string) string {
	words, space := splitSpace(text)
	if space == 0 {
		return text
	}
	if len(words) == 1 {
		res := words[0]
		for i := 0; i < space; i++ {
			res = res + " "
		}
		return res
	}
	res := ""
	for w, subStr := range words {
		res = res + subStr
		for i := 0; i < (space/(len(words)-1)) && w < len(words)-1; i++ {
			res = res + " "
		}
	}
	for i := 0; i < (space % (len(words) - 1)); i++ {
		res = res + " "
	}
	return res
}

func splitSpace(text string) ([]string, int) {
	words := make([]string, 0)
	space := 0
	start := 0
	for i, b := range text {
		if b == ' ' {
			if text[start] != ' ' {
				words = append(words, text[start:i])
				start = i
			}
			if start < len(text)-1 {
				start++
			}
			space++
		}
		if i == len(text)-1 && text[start] != ' ' {
			words = append(words, text[start:])
		}
	}
	return words, space
}
