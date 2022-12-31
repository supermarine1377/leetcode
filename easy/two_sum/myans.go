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

func twoSum_slow(nums []int, target int) []int {
	var ans [2]int
	for i, left := range nums {
		for j, right := range nums {
			if i != j && left+right == target {
				ans[0] = i
				ans[1] = j
			}
		}
	}
	var result []int
	if ans[0] > ans[1] {
		result = append(result, ans[1])
		result = append(result, ans[0])
		return result
	}
	result = append(result, ans[0])
	result = append(result, ans[1])
	return result
}
