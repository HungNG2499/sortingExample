package sortAlgorithm

func BubbleSort(arr []int) ([]int, int) {
	n := len(arr)
	comparisons := 0

	sortedArr := make([]int, n)
	copy(sortedArr, arr)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			comparisons++
			if sortedArr[j] > sortedArr[j+1] {
				sortedArr[j], sortedArr[j+1] = sortedArr[j+1], sortedArr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return sortedArr, comparisons
}
