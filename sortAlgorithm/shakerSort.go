package sortAlgorithm

func ShakerSort(arr []int) ([]int, int) {
	n := len(arr)
	comparisons := 0
	swapped := true
	start := 0
	end := n - 1

	sortedArr := make([]int, n)
	copy(sortedArr, arr)

	for swapped {
		swapped = false

		for i := start; i < end; i++ {
			comparisons++
			if sortedArr[i] > sortedArr[i+1] {
				sortedArr[i], sortedArr[i+1] = sortedArr[i+1], sortedArr[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		swapped = false
		end--

		for i := end - 1; i >= start; i-- {
			comparisons++
			if sortedArr[i] > sortedArr[i+1] {
				sortedArr[i], sortedArr[i+1] = sortedArr[i+1], sortedArr[i]
				swapped = true
			}
		}

		start++
	}

	return sortedArr, comparisons
}
