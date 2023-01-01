/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (43.59%)
 * Likes:    5657
 * Dislikes: 6434
 * Total Accepted:    1.7M
 * Total Submissions: 3.9M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * A phrase is a palindrome if, after converting all uppercase letters into
 * lowercase letters and removing all non-alphanumeric characters, it reads the
 * same forward and backward. Alphanumeric characters include letters and
 * numbers.
 *
 * Given a string s, return true if it is a palindrome, or false otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "A man, a plan, a canal: Panama"
 * Output: true
 * Explanation: "amanaplanacanalpanama" is a palindrome.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "race a car"
 * Output: false
 * Explanation: "raceacar" is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: s = " "
 * Output: true
 * Explanation: s is an empty string "" after removing non-alphanumeric
 * characters.
 * Since an empty string reads the same forward and backward, it is a
 * palindrome.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 2 * 10^5
 * s consists only of printable ASCII characters.
 *
 *
 */

package valid_palindrome

import "unicode"

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !unicode.IsLetter(rune(s[l])) {
			l++
			continue
		}
		if !unicode.IsLetter(rune(s[r])) {
			r--
			continue
		}
		sl := unicode.ToLower(rune(s[l]))
		sr := unicode.ToLower(rune(s[r]))

		if sl != sr {
			return false
		}
		l++
		r--
	}
	return true
}
