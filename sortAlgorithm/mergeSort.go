package sortAlgorithm

func MergeSort(arr []int) ([]int, int) {
	if len(arr) <= 1 {
		return arr, 0
	}

	mid := len(arr) / 2
	left, leftComparisons := MergeSort(arr[:mid])
	right, rightComparisons := MergeSort(arr[mid:])

	merged, mergeComparisons := merge(left, right)
	totalComparisons := leftComparisons + rightComparisons + mergeComparisons

	return merged, totalComparisons
}

func merge(left, right []int) ([]int, int) {
	result := make([]int, 0, len(left)+len(right))
	i, j, comparisons := 0, 0, 0

	for i < len(left) && j < len(right) {
		comparisons++
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result, comparisons
}
