package valid_perfect_square

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l, r := 0, num
	for l <= r {
		m := r + (l-r)/2
		if m*m > num {
			r = m - 1
		} else if m*m < num {
			l = m + 1
		} else {
			return true
		}
	}
	return false
}
