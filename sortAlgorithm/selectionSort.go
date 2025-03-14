package sortAlgorithm

func SelectionSort(arr []int) ([]int, int) {
	n := len(arr)
	comparisons := 0

	sortedArr := make([]int, n)
	copy(sortedArr, arr)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			comparisons++
			if sortedArr[j] < sortedArr[minIdx] {
				minIdx = j
			}
		}
		sortedArr[i], sortedArr[minIdx] = sortedArr[minIdx], sortedArr[i]
	}

	return sortedArr, comparisons
}
