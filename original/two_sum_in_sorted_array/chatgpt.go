package two_sum_in_sorted_array

func twoSum(numbers []int, target int) bool {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}
