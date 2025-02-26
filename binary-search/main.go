package main

import "fmt"

func BinarySearch(haystack []int, needle int) int {
	// make a binary search function here
	var lo = 0
	var hi = len(haystack)

	for lo < hi {
		var m = lo + (hi-lo)/2
		var v = haystack[m]

		if v == needle {
			return m
		} else if v > needle {
			hi = m
		} else if v < needle {
			lo = m + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5}, 1))
}
