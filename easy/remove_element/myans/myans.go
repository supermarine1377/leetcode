package myans

func removeElement(nums []int, val int) int {
	firstIndex := 0
	for secondIndex, num := range nums {
		if num != val {
			nums[firstIndex] = nums[secondIndex]
			firstIndex++
		}
	}
	return firstIndex
}
