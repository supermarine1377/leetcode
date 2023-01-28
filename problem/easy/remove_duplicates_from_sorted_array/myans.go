package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) []int {
	if nums == nil {
		return nil
	}
	if len(nums) == 1 {
		return nums
	}
	l, r := 0, 1
	for r < len(nums) {
		if nums[l] == nums[r] {
			r++
			continue
		}
		nums[l+1] = nums[r]
		l++
		r++
	}
	return nums[0 : l+1]
}
