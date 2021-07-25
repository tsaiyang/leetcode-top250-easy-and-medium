package main

import "math"

func maxSubArray(nums []int) int {
    cur := 0
    result := math.MinInt32
    for _, num := range nums {
        cur += num
        result = max(result, cur)
        if cur < 0 {
            cur = 0
        }
    }
    return result
}
