package binary_search

func serach(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		num := nums[m]
		if num > target {
			r = m - 1
		} else if num < target {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}
