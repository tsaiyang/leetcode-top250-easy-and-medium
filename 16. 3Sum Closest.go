package main

import (
    "math"
    "sort"
)

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    result := math.MaxInt32
    for i := range nums {
        left := i + 1
        right := len(nums) - 1
        for left < right {
            sum := nums[left] + nums[right] + nums[i]
            if sum == target {
                return target
            }
            if sum < target {
                left++
            } else {
                right--
            }
            result = min(result, abs(target-sum))
        }
    }
    return result
}
