package sortAlgorithm

func InsertionSort(arr []int) ([]int, int) {
	n := len(arr)
	comparisons := 0

	sortedArr := make([]int, n)
	copy(sortedArr, arr)

	for i := 1; i < n; i++ {
		key := sortedArr[i]
		j := i - 1

		for j >= 0 {
			comparisons++
			if sortedArr[j] > key {
				sortedArr[j+1] = sortedArr[j]
			} else {
				break
			}
			j--
		}

		sortedArr[j+1] = key
	}

	return sortedArr, comparisons
}
