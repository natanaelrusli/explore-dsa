package main

import (
	"strconv"
	"strings"
)

func StringChallenge(strArr []string) string {
	word := strArr[0]
	rowCount := strArr[1]
	challengeToken := "j1kdebf"

	resRows := make(map[int][]string)

	i := 0
	j := 0
	incr := 1
	count, _ := strconv.Atoi(rowCount)

	for i < len(word) {
		resRows[j] = append(resRows[j], string(word[i]))
		if j == (count - 1) {
			incr = -1
		} else if j == 0 {
			incr = 1
		}

		j += incr
		i++
	}

	var result string
	for i := 0; i < count; i++ {
		res := strings.Join(resRows[i], "")

		result += string(res)
	}

	// code goes here
	return intersperseString(result, challengeToken)
}
