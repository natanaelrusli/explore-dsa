package main

import "fmt"

// Time Complexity: O(n^2)
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func MergeSort(arr []int) []int {
	// preflight check
	// negative space programming
	// check if the array have no or 1 item
	if len(arr) <= 1 {
		return arr
	}

	// Step 1: Divide the array
	mid := len(arr) / 2

	// Step 2: Recursively divide each half
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	result := make([]int, 0, len(left)+len(right))
	leftIndex, rightIndex := 0, 0

	// Step 3: Sort the subarrays
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	// Step 4: Merge the final 2 halves
	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, v := range arr {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	// ... at the end means spread the elements
	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{100, 38, 27, 43, 3, 9, 82, 10}
	fmt.Println(BubbleSort(arr))
}
