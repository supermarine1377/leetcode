package reverse_only_letters

func reverseOnlyLetters(s string) string {
	runes := []rune(s)
	l, r := 0, len(s)-1
	for l < r {
		// when s[l] is not alphabetical letter
		if !((s[l] >= 65 && s[l] <= 90) || (s[l] >= 97 && s[l] <= 122)) {
			l++
			// when s[r] is not alphabetical letter
		} else if !((s[r] >= 65 && s[r] <= 90) || (s[r] >= 97 && s[r] <= 122)) {
			r--
		} else {
			runes[l], runes[r] = runes[r], runes[l]
			l++
			r--
		}
	}
	return string(runes)
}
