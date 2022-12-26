package twosum

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		wantNum := target - num
		j, ok := numMap[wantNum]
		if !ok {
			numMap[num] = i
			continue
		}
		return []int{i, j}
	}
	return nil
}
