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

func lengthOfLastWord_Arr(s string) int {
	var word []rune
	var isSp bool
	for _, r := range s {
		if string(r) == " " {
			isSp = true
			continue
		}
		if isSp && string(r) != " " {
			word = []rune{}
		}
		word = append(word, r)
		isSp = false
	}
	return len(word)
}
