package sortAlgorithm

func FlashSort(arr []int) ([]int, int) {
	n := len(arr)
	if n == 0 {
		return arr, 0
	}

	m := int(0.45 * float64(n))
	l := make([]int, m)
	min := arr[0]
	max := arr[0]
	maxIndex := 0

	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
	}

	if min == max {
		return arr, 0
	}

	c1 := (float64(m) - 1) / float64(max-min)
	for i := 0; i < n; i++ {
		k := int(c1 * float64(arr[i]-min))
		l[k]++
	}

	for i := 1; i < m; i++ {
		l[i] += l[i-1]
	}

	sortedArr := make([]int, n)
	copy(sortedArr, arr)

	sortedArr[maxIndex], sortedArr[0] = sortedArr[0], sortedArr[maxIndex]
	move := 0
	j := 0
	k := m - 1
	for move < n-1 {
		for j > l[k]-1 {
			j++
			k = int(c1 * float64(sortedArr[j]-min))
		}
		flash := sortedArr[j]
		for j != l[k] {
			k = int(c1 * float64(flash-min))
			hold := sortedArr[l[k]-1]
			sortedArr[l[k]-1] = flash
			flash = hold
			l[k]--
			move++
		}
	}

	insertionSort(sortedArr)
	return sortedArr, move
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
