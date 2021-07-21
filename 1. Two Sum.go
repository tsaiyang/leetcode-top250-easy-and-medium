package main

func twoSum(nums []int, target int) []int {
	intMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, ok := intMap[target-num]; ok {
			return []int{i, index}
		} else {
			intMap[num] = i
		}
	}
	return []int{-1, -1}
}
