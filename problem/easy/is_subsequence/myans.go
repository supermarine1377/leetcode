package is_subsequence

/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 *
 * https://leetcode.com/problems/is-subsequence/description/
 *
 * algorithms
 * Easy (49.13%)
 * Likes:    6489
 * Dislikes: 365
 * Total Accepted:    706.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '"abc"\n"ahbgdc"'
 *
 * Given two strings s and t, return true if s is a subsequence of t, or false
 * otherwise.
 *
 * A subsequence of a string is a new string that is formed from the original
 * string by deleting some (can be none) of the characters without disturbing
 * the relative positions of the remaining characters. (i.e., "ace" is a
 * subsequence of "abcde" while "aec" is not).
 *
 *
 * Example 1:
 * Input: s = "abc", t = "ahbgdc"
 * Output: true
 * Example 2:
 * Input: s = "axc", t = "ahbgdc"
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 100
 * 0 <= t.length <= 10^4
 * s and t consist only of lowercase English letters.
 *
 *
 *
 * Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k
 * >= 10^9, and you want to check one by one to see if t has its subsequence.
 * In this scenario, how would you change your code?
 */

// @lc code=start

func isSubsequence(s string, t string) bool {
	// strMap := make(map[rune]int)
	// for i, tt := range t {
	// 	strMap[tt] = i
	// }

	// var i int
	// for _, ss := range s {
	// 	index, ok := strMap[ss]
	// 	if !ok {
	// 		return false
	// 	}
	// 	if index < i {
	// 		return false
	// 	}
	// 	i = index
	// }
	// return true

	var i, j int
	for {
		if i == len(s) {
			return true
		}
		if j == len(t) {
			return false
		}
		if s[i] == t[j] {
			i++
			j++
			continue
		}
		j++
	}
}
