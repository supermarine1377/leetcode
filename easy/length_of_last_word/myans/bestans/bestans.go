package bestans

func lengthOfLastWord(s string) int {
	var count int
	for i := len(s) - 1; i > -1; i-- {
		if string(s[i]) != " " {
			count++
			continue
		}
		if count > 0 {
			return count
		}
	}
	return count
}
