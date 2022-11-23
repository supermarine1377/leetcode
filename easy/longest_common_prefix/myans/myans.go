package myans

// easy[14]
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		str := strs[i]
		prefix = findCommonPrefix(prefix, str)
	}

	return prefix
}

func findCommonPrefix(str1, str2 string) string {
	if str1 == "" || str2 == "" {
		return ""
	}
	if str1[0] != str2[0] {
		return ""
	}
	var (
		longer  string
		shorter string
	)
	if len(str1) > len(str2) {
		longer = str1
		shorter = str2
	} else {
		longer = str2
		shorter = str1
	}
	prefix := shorter
	for i := 0; i < len(shorter); i++ {
		if shorter[i] != longer[i] {
			if i != 0 {
				return shorter[0:i]
			}
			prefix = shorter[0 : i-1]
		}
	}

	return prefix
}
