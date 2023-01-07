package bestans

func plusOne(digits []int) []int {
	len := len(digits)
	for i := len - 1; i > -1; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	var a = make([]int, len+1)
	a[0] = 1
	return a
}
