package main

import (
	"fmt"
	"math"
)

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

func JumpSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	if n <= 1 {
		if arr[0] == target {
			return 0
		}
		return -1
	}

	// Finding the optimal jump step size
	step := int(math.Sqrt(float64(n)))

	// Finding the block where element is present (if exists)
	prev := 0

	// Jump to the next block as long as target is greater than current element
	for curr := step; curr < n && arr[min(curr, n-1)] < target; {
		prev = curr
		curr += step
	}

	// Do linear search within the identified block
	for j := prev; j <= min(prev+step, n); j++ {
		if arr[j] == target {
			return j
		}
		// If we find a larger element, target doesn't exist
		if arr[j] > target {
			return -1
		}
	}

	return -1
}

// Helper function to find minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example usage with different scenarios
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	sortedNumbers := []int{1, 2, 3, 4, 5}

	fmt.Println(JumpSearch(sortedNumbers, 5))

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
