package sortAlgorithm

func BinaryInsertionSort(arr []int) ([]int, int) {
	n := len(arr)
	comparisons := 0

	sortedArr := make([]int, n)
	copy(sortedArr, arr)

	for i := 1; i < n; i++ {
		key := sortedArr[i]
		left := 0
		right := i - 1

		for left <= right {
			comparisons++
			mid := left + (right-left)/2
			if sortedArr[mid] > key {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		copy(sortedArr[left+1:i+1], sortedArr[left:i])
		sortedArr[left] = key
	}

	return sortedArr, comparisons
}
