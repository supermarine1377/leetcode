package move_zeros

func moveZeroes(nums []int) {
	var i int
	for j := range nums {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
