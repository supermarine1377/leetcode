package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) []int {
	if nums == nil {
		return nil
	}
	if len(nums) == 1 {
		return nums
	}
	s, f := 0, 1
	for f < len(nums) {
		if nums[s] == nums[f] {
			f++
			continue
		}
		nums[s+1] = nums[f]
		s++
		f++
	}
	return nums[0 : s+1]
}
