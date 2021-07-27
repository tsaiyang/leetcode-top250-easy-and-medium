package main

func findPeakElement(nums []int) int {
    if len(nums) < 2 || nums[0] > nums[1] {
        return 0
    }
    if nums[len(nums)-1] > nums[len(nums)-2] {
        return len(nums) - 1
    }
    i := 1
    j := len(nums) - 2
    mid := 0
    for i < j {
        mid = (i + j) / 2
        if nums[mid+1] > nums[mid] {
            i = mid + 1
        } else {
            j = mid - 1
        }
    }
    return i
}
