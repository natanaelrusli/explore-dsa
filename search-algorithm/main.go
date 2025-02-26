package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// LinearSearch searches for target in slice arr and returns its index
// If target is not found, it returns -1
func LinearSearch(arr []int, target int) int {
	// Loop through each element in the slice
	for i := 0; i < len(arr); i++ {
		// If current element matches target, return its index
		if arr[i] == target {
			return i
		}
	}
	// If target wasn't found, return -1
	return -1
}

func main() {
	// Example usage with different scenarios
	numbers := []int{64, 34, 25, 12, 22, 11, 90}

	// Case 1: Searching for a number that exists
	result := LinearSearch(numbers, 12)
	fmt.Printf("Searching for 12: Found at index %d\n", result)

	// Case 2: Searching for a number that doesn't exist
	result = LinearSearch(numbers, 13)
	fmt.Printf("Searching for 13: Found at index %d\n", result)

	// Case 3: Searching in an empty slice
	emptySlice := []int{}
	result = LinearSearch(emptySlice, 1)
	fmt.Printf("Searching in empty slice: Found at index %d\n", result)

	numbers = make([]int, 1000)
	for i := range numbers {
		numbers[i] = i * 2 // Even numbers from 0 to 1998
	}

	// Let's compare different search scenarios
	testCases := []struct {
		target int
		desc   string
	}{
		{0, "First element"},
		{998, "Middle element"},
		{1998, "Last element"},
		{999, "Non-existing element"},
	}

	for _, tc := range testCases {
		index := BinarySearch(numbers, tc.target)
		fmt.Printf("%s: Found %d at index %d\n",
			tc.desc, tc.target, index)
	}
}
