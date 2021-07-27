package main

import "math"

func singleNumber1(nums []int) int {
    arr := make([]int, len(nums))
    copy(arr, nums)
    for i := range nums {
        if nums[i] < 0 {
            nums[i] = -nums[i]
        }
    }
    bit := make([]int, 32)
    for i := range nums {
        for j := 31; j >= 0; j-- {
            if nums[i]&1 == 1 {
                bit[j]++
            }
            nums[i] >>= 1
        }
    }
    for i := range bit {
        bit[i] %= 3
    }
    result := 0
    for i := len(bit) - 1; i >= 0; i-- {
        if bit[i] == 1 {
            result += int(math.Pow(float64(2), float64(len(bit)-1-i)))
        }
    }
    count := 0
    for i := range arr {
        if arr[i] == result {
            count++
        }
    }
    if count > 1 || count == 0 {
        return -result
    }
    return result
}
