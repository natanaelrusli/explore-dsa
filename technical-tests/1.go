package main

import (
	"strings"
)

func intersperseString(str1, str2 string) string {
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

func findOverlappingElements(arr1, arr2 []string) []string {
	result := make([]string, 0)

	arr1Map := make(map[string]bool)
	for _, char := range arr1 {
		arr1Map[char] = true
	}

	for _, char := range arr2 {
		if arr1Map[char] {
			result = append(result, char)
		}
	}

	return result
}

func ArrayChallenge(strArr []string) string {
	// code goes here
	challengeToken := "j1kdebf"

	arr1 := strings.Split(strings.ReplaceAll(strArr[0], " ", ""), ",")
	arr2 := strings.Split(strings.ReplaceAll(strArr[1], " ", ""), ",")

	result := findOverlappingElements(arr1, arr2)
	resultStr := strings.Join(result, ",")

	return intersperseString(resultStr, challengeToken)
}

// func main() {
// 	log.Println(ArrayChallenge([]string{"bob", "bob", "abc"}))
// }
