package myans

func removeElement(nums []int, val int) []int {
	i := 0
	for j := range nums {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return nums[:i]
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
