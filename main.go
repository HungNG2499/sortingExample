package main

import (
	"fmt"

	"example.com/sortingExample/sortAlgorithm"
)

func main() {

	arr := []int{4, 2, 2, 8, 3, 3, 1}
	fmt.Println("Original array:", arr)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons := sortAlgorithm.CountingSort(arr)
	fmt.Println("Sorted array (Counting Sort):", sortedArr)
	fmt.Println("Number of comparisons (Counting Sort):", comparisons)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons = sortAlgorithm.HeapSort(arr)
	fmt.Println("Sorted array (Heap Sort):", sortedArr)
	fmt.Println("Number of comparisons (Heap Sort):", comparisons)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons = sortAlgorithm.BinaryInsertionSort(arr)
	fmt.Println("Sorted array (Binary Insertion Sort):", sortedArr)
	fmt.Println("Number of comparisons (Binary Insertion Sort):", comparisons)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons = sortAlgorithm.SelectionSort(arr)
	fmt.Println("Sorted array (Selection Sort):", sortedArr)
	fmt.Println("Number of comparisons (Selection Sort):", comparisons)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons = sortAlgorithm.InsertionSort(arr)
	fmt.Println("Sorted array (Insertion Sort):", sortedArr)
	fmt.Println("Number of comparisons (Insertion Sort):", comparisons)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons = sortAlgorithm.ShellSort(arr)
	fmt.Println("Sorted array (Shell Sort):", sortedArr)
	fmt.Println("Number of comparisons (Shell Sort):", comparisons)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons = sortAlgorithm.BubbleSort(arr)
	fmt.Println("Sorted array (Bubble Sort):", sortedArr)
	fmt.Println("Number of comparisons (Bubble Sort):", comparisons)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons = sortAlgorithm.MergeSort(arr)
	fmt.Println("Sorted array (Merge Sort):", sortedArr)
	fmt.Println("Number of comparisons (Merge Sort):", comparisons)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons = sortAlgorithm.QuickSort(arr)
	fmt.Println("Sorted array (Quick Sort):", sortedArr)
	fmt.Println("Number of comparisons (Quick Sort):", comparisons)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons = sortAlgorithm.RadixSort(arr)
	fmt.Println("Sorted array (Radix Sort):", sortedArr)
	fmt.Println("Number of comparisons (Radix Sort):", comparisons)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons = sortAlgorithm.ShakerSort(arr)
	fmt.Println("Sorted array (Shaker Sort):", sortedArr)
	fmt.Println("Number of comparisons (Shaker Sort):", comparisons)
	fmt.Println("----------------------------------------")

	sortedArr, comparisons = sortAlgorithm.FlashSort(arr)
	fmt.Println("Sorted array (Flash Sort):", sortedArr)
	fmt.Println("Number of comparisons (Flash Sort):", comparisons)
	fmt.Println("----------------------------------------")
}
