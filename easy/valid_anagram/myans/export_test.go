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

import (
	"testing"
)

type args struct {
	s string
	t string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "1st",
		args: args{
			s: "a",
			t: "aa",
		},
		want: false,
	},
	{
		name: "2nd",
		args: args{
			s: "cat",
			t: "tac",
		},
		want: true,
	},
	{
		name: "3rd",
		args: args{
			s: "god",
			t: "dog",
		},
		want: true,
	},
	{
		name: "4th",
		args: args{
			s: "optp",
			t: "optt",
		},
		want: false,
	},
}

func Test_isAnagram(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagram_2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram_2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
