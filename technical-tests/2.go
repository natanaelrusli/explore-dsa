package main

import (
	"strconv"
)

func intersparseString(str1, str2 string) string {
	var result string
	len1, len2 := len(str1), len(str2)

	maxLen := len1
	if len2 > len1 {
		maxLen = len2
	}

	for i := 0; i < maxLen; i++ {
		if i < len1 {
			result += string(str1[i])
		}
		if i < len2 {
			result += string(str2[i])
		}
	}

	return result
}

func rotateArray(arr []int, positions int) []int {
	length := len(arr)
	positions = positions % length

	result := append(arr[positions:], arr[:positions]...)

	return result
}

func ArrayChallenge2(arr []int) string {

	// code goes here
	var resStr string
	challengeToken := "j1kdebf"

	resArr := rotateArray(arr, arr[0])

	for _, val := range resArr {
		resStr += strconv.Itoa(val)
	}

	result := intersparseString(resStr, challengeToken)

	return result

}
