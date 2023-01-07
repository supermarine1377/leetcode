package myans

import "strings"

func lengthOfLastWord(s string) int {
	strArr := strings.Split(s, " ")
	for i := len(strArr) - 1; i >= 0; i-- {
		elem := strArr[i]
		if elem != "" {
			return len(elem)
		}
	}
	return 0
}
