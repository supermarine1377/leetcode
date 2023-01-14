package check_the_string_is_palindrome

func isPalindrome(str string) bool {
	var st []rune
	length := len(str)
	isOdd := length%2 != 0
	for i, r := range str {
		if isOdd && i == length/2 {
			continue
		}
		if i >= length/2 {
			if st[len(st)-1] != r {
				return false
			}
			st = st[:len(st)-1]
			continue
		}
		st = append(st, r)
	}
	return true
}
