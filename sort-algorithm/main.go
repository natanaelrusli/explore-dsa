package main

import "log"

// Time Complexity: O(n^2)
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// Time Complexity: O(n log n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// divide
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// conquer
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	leftIndex, rightIndex := 0, 0

	// combine
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result
}

func main() {
	arr := []int{3, 5, 1, 3, 2, 8, 7}

	log.Println(BubbleSort(arr))
	log.Println(MergeSort(arr))
}
