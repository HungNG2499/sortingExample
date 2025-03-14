package sortAlgorithm

func ShellSort(arr []int) ([]int, int) {
	n := len(arr)
	gap := n / 2
	comparisons := 0

	sortedArr := make([]int, n)
	copy(sortedArr, arr)

	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := sortedArr[i]
			j := i
			for j >= gap {
				comparisons++
				if sortedArr[j-gap] > temp {
					sortedArr[j] = sortedArr[j-gap]
					j -= gap
				} else {
					break
				}
			}
			sortedArr[j] = temp
		}
		gap /= 2
	}

	return sortedArr, comparisons
}
