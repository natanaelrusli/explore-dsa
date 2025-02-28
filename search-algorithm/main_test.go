package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Found in middle",
			arr:      []int{1, 2, 3, 4, 5, 6, 7},
			target:   4,
			expected: 3,
		},
		{
			name:     "Found at beginning",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "Found at end",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "Not found - too small",
			arr:      []int{1, 2, 3, 4, 5},
			target:   0,
			expected: -1,
		},
		{
			name:     "Not found - too large",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "Single element - found",
			arr:      []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "Single element - not found",
			arr:      []int{1},
			target:   2,
			expected: -1,
		},
		{
			name:     "Duplicate elements",
			arr:      []int{1, 2, 2, 2, 3},
			target:   2,
			expected: 2, // Returns one of the positions where 2 appears
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("BinarySearch() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Found in middle",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "Found at beginning",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "Found at end",
			arr:      []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
		{
			name:     "Not found",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "Single element - found",
			arr:      []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "Single element - not found",
			arr:      []int{1},
			target:   2,
			expected: -1,
		},
		{
			name:     "Duplicate elements - returns first occurrence",
			arr:      []int{1, 2, 2, 2, 3},
			target:   2,
			expected: 1,
		},
		{
			name:     "Unsorted array",
			arr:      []int{5, 2, 8, 1, 9},
			target:   8,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LinearSearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("LinearSearch() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Benchmark functions
func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

func generateRandomArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size * 2) // Multiply by 2 to spread out the numbers
	}
	return arr
}

func BenchmarkBinarySearch(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			arr := generateSortedArray(size)
			target := size / 2 // Search for middle element
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				BinarySearch(arr, target)
			}
		})
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			arr := generateRandomArray(size)
			target := arr[size/2] // Search for middle element
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				LinearSearch(arr, target)
			}
		})
	}
}

// TestSearchComparison tests both search algorithms with the same input
func TestSearchComparison(t *testing.T) {
	// Generate a sorted array for binary search
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 13
	expectedIndex := 6

	// Test both search algorithms
	binaryResult := BinarySearch(arr, target)
	linearResult := LinearSearch(arr, target)

	if binaryResult != expectedIndex {
		t.Errorf("BinarySearch() = %v, want %v", binaryResult, expectedIndex)
	}

	if linearResult != expectedIndex {
		t.Errorf("LinearSearch() = %v, want %v", linearResult, expectedIndex)
	}

	// Verify both algorithms return the same result
	if binaryResult != linearResult {
		t.Errorf("Search results don't match: Binary=%v, Linear=%v", binaryResult, linearResult)
	}
}
