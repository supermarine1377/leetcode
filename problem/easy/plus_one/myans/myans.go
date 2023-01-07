package myans

func plusOne(digits []int) []int {
	len := len(digits)
	if digits[len-1] != 9 {
		digits[len-1] = digits[len-1] + 1
		return digits
	}
	for i := len - 1; i > -1; i-- {
		digit := digits[i]
		if digit == 9 {
			digits[i] = 0
			continue
		}
		digits[i] = digit + 1
		return digits
	}
	newDigits := make([]int, len+1)
	newDigits[0] = 1
	return newDigits
}
