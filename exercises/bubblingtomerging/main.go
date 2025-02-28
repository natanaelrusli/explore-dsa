package main

import (
	"fmt"
	"log"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func mergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	mid := n / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	result := make([]int, 0, len(left)+len(right))
	leftIndex, rightIndex := 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	log.Println(left)
	log.Println(right)

	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	// fmt.Println("Unsorted array:", arr)
	// bubbleSort(arr)
	// fmt.Println("Sorted array:", arr)

	fmt.Println("Unsorted array:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
