package main

func search(nums []int, target int) int {
    i := 0
    j := len(nums) - 1
    mid := 0
    for i+1 < j {
        mid = (i + j) / 2
        if target == nums[mid] {
            return mid
        }
        if nums[mid] > nums[i] {
            if target >= nums[i] && target < nums[mid] {
                j = mid - 1
            } else {
                i = mid + 1
            }
        } else {
            if target <= nums[j] && target > nums[mid] {
                i = mid + 1
            } else {
                j = mid - 1
            }
        }
    }
    if nums[i] == target {
        return i
    }
    if nums[j] == target {
        return j
    }
    return -1
}
