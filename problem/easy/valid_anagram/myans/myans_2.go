package myans

// Time complexity is O(N), but space complexity is higer than the sorting solution
func isAnagram_2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapForS := make(map[rune]int)
	for _, ss := range s {
		_, ok := mapForS[ss]
		if !ok {
			mapForS[ss] = 1
		}
		mapForS[ss] = mapForS[ss] + 1
	}

	mapForT := make(map[rune]int)
	for _, tt := range t {
		_, ok := mapForT[tt]
		if !ok {
			mapForT[tt] = 1
		}
		mapForT[tt] = mapForT[tt] + 1
	}

	for ss, numS := range mapForS {
		numT, ok := mapForT[ss]
		if !ok {
			return false
		}
		if numS != numT {
			return false
		}
	}

	return true
}

// @lc code=end
