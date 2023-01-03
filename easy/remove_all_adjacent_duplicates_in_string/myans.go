/*
 * @lc app=leetcode id=1047 lang=golang
 *
 * [1047] Remove All Adjacent Duplicates In String
 *
 * https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/description/
 *
 * algorithms
 * Easy (70.47%)
 * Likes:    5281
 * Dislikes: 206
 * Total Accepted:    414.1K
 * Total Submissions: 590.4K
 * Testcase Example:  '"abbaca"'
 *
 * You are given a string s consisting of lowercase English letters. A
 * duplicate removal consists of choosing two adjacent and equal letters and
 * removing them.
 *
 * We repeatedly make duplicate removals on s until we no longer can.
 *
 * Return the final string after all such duplicate removals have been made. It
 * can be proven that the answer is unique.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "abbaca"
 * Output: "ca"
 * Explanation:
 * For example, in "abbaca" we could remove "bb" since the letters are adjacent
 * and equal, and this is the only possible move.  The result of this move is
 * that the string is "aaca", of which only "aa" is possible, so the final
 * string is "ca".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "azxxzy"
 * Output: "ay"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^5
 * s consists of lowercase English letters.
 *
 *
 */

package remove_all_adjacent_duplicates_in_string

func removeDuplicates_ON3(s string) string {
	if len(s) <= 1 {
		return s
	}
	for {
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				s = removeDuplicate(s)
				continue
			}
		}
		break
	}
	return s
}

func removeDuplicate(s string) string {
	var stack string
	for _, r := range s {
		if len(stack) == 0 {
			stack = stack + string(r)
			continue
		}
		// case: len(stack) > 0
		if stack[len(stack)-1] == byte(r) {
			stack = stack[0 : len(stack)-1]
			continue
		}
		stack = stack + string(r)
	}
	return stack
}
