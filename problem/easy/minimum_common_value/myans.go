package minimum_common_value

// Use Map; time complexity is O(N+M) and space complexity is O(N)
func getCommon(nums1 []int, nums2 []int) int {
	m := make(map[int]struct{})
	result := -1
	// O(N)
	for _, num1 := range nums1 {
		m[num1] = struct{}{}
	}
	// O(M)
	for _, num2 := range nums2 {
		if _, ok := m[num2]; ok {
			if result == -1 {
				result = num2
				continue
			}
			if result > num2 {
				result = num2
			}
		}
	}
	return result
}

func getCommonTwoPointers(nums1 []int, nums2 []int) int {
	var i, j int
	result := -1
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return result
}
