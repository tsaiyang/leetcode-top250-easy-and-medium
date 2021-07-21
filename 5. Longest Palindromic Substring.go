package main

func longestPalindrome(s string) string {
	flag := make([][]bool, len(s))
	for i := range flag {
		flag[i] = make([]bool, len(s))
	}
	result := ""
	for j := 0; j < len(s); j++ {
		for i := j; i > -1; i-- {
			if s[i] == s[j] && (j-i < 2 || flag[i+1][j-1]) {
				flag[i][j] = true
				if j-i+1 > len(result) {
					result = s[i : j+1]
				}
			}
		}
	}
	return result
}
