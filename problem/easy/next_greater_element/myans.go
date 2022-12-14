/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 *
 * https://leetcode.com/problems/next-greater-element-i/description/
 *
 * algorithms
 * Easy (71.34%)
 * Likes:    5213
 * Dislikes: 324
 * Total Accepted:    474K
 * Total Submissions: 663.4K
 * Testcase Example:  '[4,1,2]\n[1,3,4,2]'
 *
 * The next greater element of some element x in an array is the first greater
 * element that is to the right of x in the same array.
 *
 * You are given two distinct 0-indexed integer arrays nums1 and nums2, where
 * nums1 is a subset of nums2.
 *
 * For each 0 <= i < nums1.length, find the index j such that nums1[i] ==
 * nums2[j] and determine the next greater element of nums2[j] in nums2. If
 * there is no next greater element, then the answer for this query is -1.
 *
 * Return an array ans of length nums1.length such that ans[i] is the next
 * greater element as described above.
 *
 *
 * Example 1:
 *
 *
 * Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
 * Output: [-1,3,-1]
 * Explanation: The next greater element for each value of nums1 is as follows:
 * - 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so
 * the answer is -1.
 * - 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.
 * - 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so
 * the answer is -1.
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [2,4], nums2 = [1,2,3,4]
 * Output: [3,-1]
 * Explanation: The next greater element for each value of nums1 is as follows:
 * - 2 is underlined in nums2 = [1,2,3,4]. The next greater element is 3.
 * - 4 is underlined in nums2 = [1,2,3,4]. There is no next greater element, so
 * the answer is -1.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums1.length <= nums2.length <= 1000
 * 0 <= nums1[i], nums2[i] <= 10^4
 * All integers in nums1 and nums2 are unique.
 * All the integers of nums1 also appear in nums2.
 *
 *
 *
 * Follow up: Could you find an O(nums1.length + nums2.length) solution?
 */

package next_greater_element

func nextGreaterElement_O2(nums1 []int, nums2 []int) []int {
	nums2Map := make(map[int]int)
	for i, num2 := range nums2 {
		nums2Map[num2] = i
	}
	var result []int
	for _, num1 := range nums1 {
		i := nums2Map[num1]
		newNums2 := nums2[i:]
		var firstGreater int
		for _, num2 := range newNums2 {
			if num2 > num1 {
				firstGreater = num2
				break
			}
		}
		if firstGreater == 0 {
			result = append(result, -1)
			continue
		}
		result = append(result, firstGreater)
	}
	return result
}

// Key observation:
// Suppose we have a decreasing sequence followed by a greater number
// For example [5, 4, 3, 2, 1, 6] then the greater number 6 is the next greater element for all previous numbers in the sequence

// We use a stack to keep a decreasing sub-sequence, whenever we see a number x greater than stack.peek() we pop all elements less than x and for all the popped ones, their next greater element is x
// For example [9, 8, 7, 3, 2, 1, 6]
// The stack will first contain [9, 8, 7, 3, 2, 1] and then we see 6 which is greater than 1 so we pop 1 2 3 whose next greater element should be 6
func nextGreaterElement_O_LenNum1_LenNum2(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int)
	var stack []int
	for _, num := range nums2 {
		for len(stack) != 0 && stack[len(stack)-1] < num {
			numMap[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	var reuslts []int
	for i := 0; i < len(nums1); i++ {
		numInMap, ok := numMap[nums1[i]]
		if !ok {
			reuslts = append(reuslts, -1)
			continue
		}
		reuslts = append(reuslts, numInMap)
	}
	return reuslts
}
