package myans

import "sort"

// @lc code=start
func searchInsert(nums []int, target int) int {
	nums = append(nums, target)
	sort.Ints(nums)

	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return 0
}
