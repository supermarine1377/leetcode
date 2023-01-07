package replace_elements_with_greatest_element_on_right_side

/*
 * @lc app=leetcode id=1299 lang=golang
 *
 * [1299] Replace Elements with Greatest Element on Right Side
 *
 * https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/description/
 *
 * algorithms
 * Easy (74.73%)
 * Likes:    1740
 * Dislikes: 180
 * Total Accepted:    261.5K
 * Total Submissions: 351.5K
 * Testcase Example:  '[17,18,5,4,6,1]'
 *
 * Given an array arr, replace every element in that array with the greatest
 * element among the elements to its right, and replace the last element with
 * -1.
 *
 * After doing so, return the array.
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [17,18,5,4,6,1]
 * Output: [18,6,6,6,1,-1]
 * Explanation:
 * - index 0 --> the greatest element to the right of index 0 is index 1 (18).
 * - index 1 --> the greatest element to the right of index 1 is index 4 (6).
 * - index 2 --> the greatest element to the right of index 2 is index 4 (6).
 * - index 3 --> the greatest element to the right of index 3 is index 4 (6).
 * - index 4 --> the greatest element to the right of index 4 is index 5 (1).
 * - index 5 --> there are no elements to the right of index 5, so we put -1.
 *
 *
 * Example 2:
 *
 *
 * Input: arr = [400]
 * Output: [-1]
 * Explanation: There are no elements to the right of index 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 10^4
 * 1 <= arr[i] <= 10^5
 *
 *
 */

// @lc code=start

func replaceElementsO2(arr []int) []int {
	if len(arr) < 2 {
		return []int{-1}
	}
	for i := 0; i < len(arr)-1; i++ {
		newArr := arr[i+1:]
		max := newArr[0]
		for j := 1; j < len(newArr); j++ {
			if newArr[j] > max {
				max = newArr[j]
			}
		}
		arr[i] = max
	}
	arr[len(arr)-1] = -1
	return arr
}

func replaceElementsO1(arr []int) []int {
	if len(arr) < 2 {
		return []int{-1}
	}
	result := make([]int, len(arr))
	max := -1
	result[len(result)-1] = max

	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > max {
			max = arr[i]
		}
		result[i-1] = max
	}
	return result
}
