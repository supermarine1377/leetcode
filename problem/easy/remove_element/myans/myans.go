package myans

func removeElement(nums []int, val int) []int {
	firstIndex := 0
	for secondIndex := range nums {
		if nums[secondIndex] != val {
			nums[firstIndex] = nums[secondIndex]
			firstIndex++
		}
	}
	return nums[:firstIndex]
}

// This solution does repeted-work
func removeElement_notUsingTwoPointers(nums []int, val int) []int {
	if len(nums) == 0 {
		return nil
	}
	var i int
	for {
		if nums[i] == val {
			// remove element
			nums = append(nums[:i], nums[i+1:]...)
			if len(nums) == 0 {
				return []int{}
			}
			i = 0
			continue
		}
		if i == len(nums)-1 {
			break
		}
		i++
	}
	return nums
}
