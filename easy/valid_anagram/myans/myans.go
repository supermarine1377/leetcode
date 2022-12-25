/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (62.74%)
 * Likes:    7763
 * Dislikes: 249
 * Total Accepted:    1.9M
 * Total Submissions: 3M
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t, return true if t is an anagram of s, and false
 * otherwise.
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a
 * different word or phrase, typically using all the original letters exactly
 * once.
 *
 *
 * Example 1:
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 * Example 2:
 * Input: s = "rat", t = "car"
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length, t.length <= 5 * 10^4
 * s and t consist of lowercase English letters.
 *
 *
 *
 * Follow up: What if the inputs contain Unicode characters? How would you
 * adapt your solution to such a case?
 *
 */

// @lc code=starts
package myans

import "sort"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var arr1 []string
	for i := 0; i < len(s); i++ {
		arr1 = append(arr1, string(s[i]))
	}
	sort.Strings(arr1)

	var arr2 []string
	for i := 0; i < len(t); i++ {
		arr2 = append(arr2, string(t[i]))
	}
	sort.Strings(arr2)

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// @lc code=end
