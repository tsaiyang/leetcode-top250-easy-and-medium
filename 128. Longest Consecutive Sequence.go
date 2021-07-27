package main

func longestConsecutive(nums []int) int {
    numMap := make(map[int]int)
    result := 0
    for i := range nums {
        if _, ok := numMap[nums[i]]; !ok {
            numMap[nums[i]] = 1
            if _, ok := numMap[nums[i]-1]; ok {
                result = max(result, merge2(numMap, nums[i]-1, nums[i]))
            }
            if _, ok := numMap[nums[i]+1]; ok {
                result = max(result, merge2(numMap, nums[i], nums[i]+1))
            }
        }
    }
    return result
}

func merge2(numMap map[int]int, less, more int) int {
    left := numMap[less] - less + 1
    right := numMap[more] + more - 1
    length := right - left + 1
    numMap[left] = length
    numMap[right] = length
    return length
}
