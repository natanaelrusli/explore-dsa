package main

import "fmt"

func countSetBits(n int) int {
	// Write your code here
	count := 0

	for n > 0 {
		count += n & 1
		n >>= 1
	}

	return count
}

func main() {
	num := 42                      // Binary: 101010
	fmt.Println(countSetBits(num)) // Should print 3
}
