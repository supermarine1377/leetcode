package myans

import "math"

func plusOne(digits []int) []int {
	num := convertToInt(digits)
	num++
	return convertToArr(num)
}

// [1, 2, 3] gives 123
func convertToInt(nums []int) int {
	var result int
	for i, num := range nums {
		result += num * pow10((len(nums) - i - 1))
	}

	return result
}

// 123 gives [1, 2, 3]
func convertToArr(num int) []int {
	digit := digit(num)
	result := make([]int, digit, digit)
	for i := digit - 1; i > -1; i-- {
		rem := num % 10
		result[i] = rem
		num = num / 10
	}
	return result
}

func pow10(num int) int {
	return int(math.Pow10(num))
}

func digit(num int) int {
	var result int
	for {
		if pow10(result) <= num && pow10(result+1) > num {
			break
		}
		result++
	}
	return result + 1
}
