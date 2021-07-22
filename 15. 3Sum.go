package main

import "sort"

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var result [][]int
    for i := range nums {
        if i == 0 || nums[i] != nums[i-1] {
            left := i + 1
            right := len(nums) - 1
            for left < right {
                if nums[left]+nums[right] == -nums[i] && (left == i+1 || nums[left] != nums[left-1]) {
                    result = append(result, []int{nums[i], nums[left], nums[right]})
                    left++
                    right--
                } else if nums[left]+nums[right] > -nums[i] {
                    right--
                } else {
                    left++
                }
            }
        }
    }
    return result
}
