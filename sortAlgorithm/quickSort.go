package sortAlgorithm

func QuickSort(arr []int) ([]int, int) {
	if len(arr) < 2 {
		return arr, 0
	}

	pivot := arr[len(arr)/2]
	var leftPart, rightPart []int
	var middle []int
	comparisons := len(arr) - 1

	for _, num := range arr {
		if num < pivot {
			leftPart = append(leftPart, num)
		} else if num > pivot {
			rightPart = append(rightPart, num)
		} else {
			middle = append(middle, num)
		}
	}

	sortedLeft, leftComparisons := QuickSort(leftPart)
	sortedRight, rightComparisons := QuickSort(rightPart)

	return append(append(sortedLeft, middle...), sortedRight...), comparisons + leftComparisons + rightComparisons
}
