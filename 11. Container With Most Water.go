package main

func maxArea(height []int) int {
	result := 0
	i, j := 0, len(height)-1
	for i < j {
		result = max(result, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j++
		}
	}
	return result
}
