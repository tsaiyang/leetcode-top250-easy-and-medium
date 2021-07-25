package main

func jump(nums []int) int {
    steps := 0
    end := 0
    maxPos := 0
    for i := range nums {
        if i == end {
            steps++
            end = maxPos
        }
        maxPos = max(maxPos, nums[i]+i)
    }
    return steps
}
