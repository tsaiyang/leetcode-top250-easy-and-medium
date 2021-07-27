package main

func findMin1(nums []int) int {
    i := 0
    j := len(nums) - 1
    mid := 0
    for i+1 < j {
        mid = (i + j) / 2
        if nums[mid] > nums[i] {
            i = mid + 1
        } else {
            j = mid - 1
        }
    }
    if nums[i] < nums[j] {
        return i
    }
    return j
}
