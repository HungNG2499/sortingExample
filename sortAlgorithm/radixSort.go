package sortAlgorithm

func RadixSort(arr []int) ([]int, int) {
	if len(arr) == 0 {
		return arr, 0
	}

	max := arr[0]
	comparisons := 0

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	exp := 1
	for max/exp > 0 {
		var digitComparisons int
		arr, digitComparisons = countingSortByDigit(arr, exp)
		comparisons += digitComparisons
		exp *= 10
	}

	return arr, comparisons
}

func countingSortByDigit(arr []int, exp int) ([]int, int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)
	comparisons := 0

	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
		comparisons++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}

	return arr, comparisons
}
