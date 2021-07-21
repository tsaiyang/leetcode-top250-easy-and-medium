package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	base := 1
	for x/base >= 10 {
		base *= 10
	}
	for x != 0 {
		if x%10 != x/base {
			return false
		}
		x = (x % base) / 10
		base /= 100
	}
	return true
}
