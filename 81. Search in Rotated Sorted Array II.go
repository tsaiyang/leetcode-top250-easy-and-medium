package main

func search2(nums []int, target int) bool {
    i := 0
    j := len(nums) - 1
    mid := 0
    for i <= j {
        mid = (i + j) / 2
        if nums[mid] == target {
            return true
        }

        if nums[i] == nums[mid] && nums[mid] == nums[j] {
            for i != mid && nums[i] == nums[mid] {
                i++
            }
            if i == mid {
                i = mid + 1
                continue
            }
        }

        if nums[i] != nums[mid] {
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
        } else {
            if nums[mid] < nums[j] {
                if target <= nums[j] && target > nums[mid] {
                    i = mid + 1
                } else {
                    j = mid - 1
                }
            } else {
                if target < nums[mid] && target >= nums[i] {
                    j = mid - 1
                } else {
                    i = mid + 1
                }
            }
        }
    }
    return false
}
