package main

import (
    "reflect"
    "testing"
)

func TestBubbleSort(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected []int
    }{
        {
            name:     "Regular array",
            input:    []int{64, 34, 25, 12, 22, 11, 90},
            expected: []int{11, 12, 22, 25, 34, 64, 90},
        },
        {
            name:     "Array with duplicates",
            input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
            expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
        },
        {
            name:     "Already sorted array",
            input:    []int{1, 2, 3, 4, 5},
            expected: []int{1, 2, 3, 4, 5},
        },
        {
            name:     "Reverse sorted array",
            input:    []int{5, 4, 3, 2, 1},
            expected: []int{1, 2, 3, 4, 5},
        },
        {
            name:     "Empty array",
            input:    []int{},
            expected: []int{},
        },
        {
            name:     "Single element array",
            input:    []int{1},
            expected: []int{1},
        },
        {
            name:     "Negative numbers",
            input:    []int{-3, -5, 0, 1, -4, 2, -1},
            expected: []int{-5, -4, -3, -1, 0, 1, 2},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create a copy of input to avoid modifying the original
            input := make([]int, len(tt.input))
            copy(input, tt.input)

            result := BubbleSort(input)

            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("BubbleSort() = %v, want %v", result, tt.expected)
            }
        })
    }
}

func TestMergeSort(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected []int
    }{
        {
            name:     "Regular array",
            input:    []int{64, 34, 25, 12, 22, 11, 90},
            expected: []int{11, 12, 22, 25, 34, 64, 90},
        },
        {
            name:     "Array with duplicates",
            input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
            expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
        },
        {
            name:     "Already sorted array",
            input:    []int{1, 2, 3, 4, 5},
            expected: []int{1, 2, 3, 4, 5},
        },
        {
            name:     "Reverse sorted array",
            input:    []int{5, 4, 3, 2, 1},
            expected: []int{1, 2, 3, 4, 5},
        },
        {
            name:     "Empty array",
            input:    []int{},
            expected: []int{},
        },
        {
            name:     "Single element array",
            input:    []int{1},
            expected: []int{1},
        },
        {
            name:     "Negative numbers",
            input:    []int{-3, -5, 0, 1, -4, 2, -1},
            expected: []int{-5, -4, -3, -1, 0, 1, 2},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Create a copy of input to avoid modifying the original
            input := make([]int, len(tt.input))
            copy(input, tt.input)

            result := MergeSort(input)

            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("MergeSort() = %v, want %v", result, tt.expected)
            }
        })
    }
}

// Benchmark functions to compare performance
func BenchmarkBubbleSort(b *testing.B) {
    input := []int{64, 34, 25, 12, 22, 11, 90}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        // Create a copy of input for each iteration
        testInput := make([]int, len(input))
        copy(testInput, input)
        BubbleSort(testInput)
    }
}

func BenchmarkMergeSort(b *testing.B) {
    input := []int{64, 34, 25, 12, 22, 11, 90}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        // Create a copy of input for each iteration
        testInput := make([]int, len(input))
        copy(testInput, input)
        MergeSort(testInput)
    }
}

// Helper function to generate large random arrays for benchmarking
func generateLargeArray(size int) []int {
    arr := make([]int, size)
    for i := 0; i < size; i++ {
        arr[i] = i
    }
    // Shuffle the array
    for i := len(arr) - 1; i > 0; i-- {
        j := i - 1
        arr[i], arr[j] = arr[j], arr[i]
    }
    return arr
}

// Benchmarks with larger arrays
func BenchmarkBubbleSort1000(b *testing.B) {
    input := generateLargeArray(1000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        testInput := make([]int, len(input))
        copy(testInput, input)
        BubbleSort(testInput)
    }
}

func BenchmarkMergeSort1000(b *testing.B) {
    input := generateLargeArray(1000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        testInput := make([]int, len(input))
        copy(testInput, input)
        MergeSort(testInput)
    }
}