package main

func removeDuplicates(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    slow := 1
    fast := 1
    for fast < len(nums) {
        if nums[fast] == nums[slow-1] {
            fast++
            continue
        }
        nums[slow] = nums[fast]
        slow++
    }
    return slow
}
