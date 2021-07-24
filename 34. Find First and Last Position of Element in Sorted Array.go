package main

func searchRange(nums []int, target int) []int {
    result := make([]int, 2)
    result[0] = firstPos(nums, target)
    if result[0] == -1 {
        result[1] = -1
    } else {
        result[1] = lastPos(nums, target)
    }
    return result
}

func firstPos(nums []int, target int) int {
    i := 0
    j := len(nums) - 1
    mid := 0
    for i+1 < j {
        mid = (i + j) / 2
        if target == nums[mid] {
            j = mid
        } else if target < nums[mid] {
            j = mid - 1
        } else {
            i = mid + 1
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

func lastPos(nums []int, target int) int {
    i := 0
    j := len(nums) - 1
    mid := 0
    for i+1 < j {
        mid = (i + j) / 2
        if target == nums[mid] {
            i = mid
        } else if target < nums[mid] {
            j = mid - 1
        } else {
            i = mid + 1
        }
    }
    if nums[j] == target {
        return j
    }
    if nums[i] == target {
        return i
    }
    return -1
}
