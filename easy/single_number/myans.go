/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 *
 * https://leetcode.com/problems/single-number/description/
 *
 * algorithms
 * Easy (70.06%)
 * Likes:    12476
 * Dislikes: 475
 * Total Accepted:    2M
 * Total Submissions: 2.8M
 * Testcase Example:  '[2,2,1]'
 *
 * Given a non-empty array of integers nums, every element appears twice except
 * for one. Find that single one.
 *
 * You must implement a solution with a linear runtime complexity and use only
 * constant extra space.
 *
 *
 * Example 1:
 * Input: nums = [2,2,1]
 * Output: 1
 * Example 2:
 * Input: nums = [4,1,2,1,2]
 * Output: 4
 * Example 3:
 * Input: nums = [1]
 * Output: 1
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -3 * 10^4 <= nums[i] <= 3 * 10^4
 * Each element in the array appears twice except for one element which appears
 * only once.
 *
 *
 */

// @lc code=start

package single_number

func singleNumber(nums []int) int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		_, ok := freqMap[num]
		if !ok {
			freqMap[num] = 1
			continue
		}
		freqMap[num] = 2
	}
	for key, freq := range freqMap {
		if freq == 1 {
			return key
		}
	}
	return 0
}
