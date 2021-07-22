package main

import "sort"

func fourSum(nums []int, target int) [][]int {
    if len(nums) < 4 {
        return [][]int{}
    }
    sort.Ints(nums)
    result := make([][]int, 0)
    for i := 0; i < len(nums)-3; i++ {
        if i == 0 || nums[i] != nums[i-1] {
            for j := i + 1; j < len(nums)-2; j++ {
                if j == i+1 || nums[j] != nums[j-1] {
                    left := j + 1
                    right := len(nums) - 1
                    for left < right {
                        if nums[i]+nums[j]+nums[left]+nums[right] == target && (left == j+1 || nums[left] != nums[left-1]) {
                            result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
                            left++
                            right--
                        } else if nums[i]+nums[j]+nums[left]+nums[right] < target {
                            left++
                        } else {
                            right--
                        }
                    }
                }
            }
        }
    }
    return result
}
