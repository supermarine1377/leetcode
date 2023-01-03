package backspace_string_compare

// if you use the fuction output() below, Time Complexity is O(N), and Space Complexity is O(len(s) + len(t))
func backspaceCompare(s string, t string) bool {
	return output(s) == output(t)
}

// Time Complexity: O(N), Space Complexity: O(len(s))
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
