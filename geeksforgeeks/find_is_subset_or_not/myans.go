package find_is_subset_or_not

// Given two arrays: arr1[0..m-1] and arr2[0..n-1].
// Find whether arr2[] is a subset of arr1[] or not.
// Both arrays are not in sorted order.
// It may be assumed that elements in both arrays are distinct.

// Time Complexity: O(N+N), Space Complexity: O(N)
func isSubset(arr1, arr2 []int) bool {
	iMap := make(map[int]int)
	for _, num := range arr1 {
		iMap[num] = 0
	}
	for _, num := range arr2 {
		_, ok := iMap[num]
		if !ok {
			return false
		}
	}
	return true
}
