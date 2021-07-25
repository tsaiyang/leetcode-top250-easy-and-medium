package main

func canJump(nums []int) bool {
    maxPos := 0
    for i := range nums {
        if i > maxPos {
            return false
        }
        maxPos = max(maxPos, i+nums[i])
    }
    return true
}
