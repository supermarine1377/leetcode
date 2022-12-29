package myans

/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (61.28%)
 * Likes:    7659
 * Dislikes: 1071
 * Total Accepted:    2.3M
 * Total Submissions: 3.8M
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an integer array nums, return true if any value appears at least twice
 * in the array, and return false if every element is distinct.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3,1]
 * Output: true
 * Example 2:
 * Input: nums = [1,2,3,4]
 * Output: false
 * Example 3:
 * Input: nums = [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

import "sort"

// @lc code=start
func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func containsDuplicateO1(nums []int) bool {
	numMap := make(map[int]interface{})
	for _, num := range nums {
		isDuplicate := numMap[num]
		if isDuplicate != nil {
			return true
		}
		numMap[num] = struct{}{}
	}
	return false
}

// @lc code=end
