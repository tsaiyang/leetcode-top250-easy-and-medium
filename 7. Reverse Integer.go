package main

import "math"

func reverse(x int) int {
	isPositive := false
	if x > 0 {
		x = -x
		isPositive = true
	}
	result := 0
	for x != 0 {
		n := x % 10
		result = result*10 + n
		if result < math.MinInt32/10 || (result == math.MaxInt32/10 && n < math.MinInt32%10) {
			return 0
		}
		x = x / 10
	}
	if isPositive {
		result = -result
	}
	return result
}
