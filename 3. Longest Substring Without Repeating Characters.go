package main

func lengthOfLongestSubstring(s string) int {
	result, slow, fast := 0, 0, 0
	charMap := make([]bool, 256)
	for fast < len(s) {
		if !charMap[s[fast]] {
			charMap[s[fast]] = true
			fast++
		} else {
			result = max(result, fast-slow)
			for s[slow] != s[fast] {
				charMap[s[slow]] = false
				slow++
			}
			charMap[s[slow]] = false
			slow++
		}
	}
	result = max(result, fast-slow)
	return result
}
