package sortAlgorithm

func CountingSort(arr []int) ([]int, int) {
	if len(arr) == 0 {
		return arr, 0
	}

	min, max := arr[0], arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	count := make([]int, max-min+1)
	for _, num := range arr {
		count[num-min]++
	}

	sortedArr := make([]int, 0, len(arr))
	comparisons := 0
	for i, c := range count {
		for c > 0 {
			sortedArr = append(sortedArr, i+min)
			c--
			comparisons++
		}
	}

	return sortedArr, comparisons
}
