package main

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) < 1 {
		return 0
	}
	//if len(s) == 1 && (s[0] < '0' || s[0] > '9') {
	//	return 0
	//}
	//if s[0] != '-' && (s[0] < '0' || s[0] > '9') {
	//	return 0
	//}
	//if s[0] == '-' && (s[1] < '0' || s[1] > '9') {
	//	return 0
	//}

	result := 0
	isPositive := true
	index := 0
	if s[0] == '-' || s[0] == '+' {
		index = 1
		if s[0] == '-' {
			isPositive = false
		}
	}
	for ; index < len(s); index++ {
		if s[index] > '9' || s[index] < '0' {
			break
		}
		num := - int(s[index] - '0')
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && num < math.MinInt32%10) {
			if isPositive{
				return math.MaxInt32
			}
			return math.MinInt32
		}
		result = result*10 + num
	}
	if isPositive {
		result = -result
	}
	if result > math.MaxInt32{
		return math.MaxInt32
	}
	return result
}
