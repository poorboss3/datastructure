package sort

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

/**
a b c
(a-b)(b-c)>0 a>b,b>c or a<b,b<c, b
(a-c)(c-b)>0 a>c,c>b or a<c,c<b, c
**/
