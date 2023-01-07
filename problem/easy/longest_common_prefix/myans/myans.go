// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Constraints:

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters.

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

func longestCommonPrefixO2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var n int
	var result string
	for {
		if len(strs[0]) == 0 {
			return ""
		}
		if n > len(strs[0])-1 {
			return result
		}
		p := strs[0][n]
		for _, str := range strs {
			if len(str) < n+1 {
				return result
			}
			if str[n] != p {
				return result
			}
		}
		result = result + string(p)
		n++
	}
}
