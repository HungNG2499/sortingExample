package sortAlgorithm

func HeapSort(arr []int) ([]int, int) {
	n := len(arr)
	comparisons := 0

	sortedArr := make([]int, n)
	copy(sortedArr, arr)

	for i := n/2 - 1; i >= 0; i-- {
		comparisons += heapify(sortedArr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		sortedArr[0], sortedArr[i] = sortedArr[i], sortedArr[0]

		comparisons += heapify(sortedArr, i, 0)
	}

	return sortedArr, comparisons
}

func heapify(arr []int, n int, i int) int {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	comparisons := 0

	if left < n {
		comparisons++
		if arr[left] > arr[largest] {
			largest = left
		}
	}

	if right < n {
		comparisons++
		if arr[right] > arr[largest] {
			largest = right
		}
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		comparisons += heapify(arr, n, largest)
	}

	return comparisons
}
