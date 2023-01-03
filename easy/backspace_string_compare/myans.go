package backspace_string_compare

func backspaceCompare(s string, t string) bool {
	return output(s) == output(t)
}

func output(s string) string {
	var result string
	for _, r := range s {
		if string(r) != "#" {
			result = result + string(r)
			continue
		}
		// case: string(r) == "#"
		if result == "" {
			continue
		}
		result = result[0 : len(result)-1]
	}
	return result
}
