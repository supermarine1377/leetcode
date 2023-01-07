package myans

func removeElement(nums []int, Val int) int {
	firstIndex := 0
	for secondIndex, num := range nums {
		if num != Val {
			nums[firstIndex] = nums[secondIndex]
			firstIndex++
		}
	}
	return firstIndex
}
