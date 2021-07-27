package main

func maxProduct(nums []int) int {
    mins := make([]int, len(nums))
    maxs := make([]int, len(nums))
    result := nums[0]
    mins[0] = nums[0]
    maxs[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        result = max(maxs[i-1]*nums[i], mins[i-1]*nums[i])
        maxs[i] = max(result, max(nums[i], 0))
        mins[i] = min(result, nums[i])
    }
    return result
}
