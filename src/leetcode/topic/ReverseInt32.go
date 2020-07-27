package topic

import "math"

func reverse(x int) int {
	result := 0
	for x != 0 {
		remainder := x % 10
		x /= 10
		result = result*10 + remainder
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	result := 0
	innerX := x
	for innerX != 0 {
		remainder := innerX % 10
		innerX /= 10
		result = result*10 + remainder
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return false
	}
	return result == x
}
