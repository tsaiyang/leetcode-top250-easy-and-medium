package main

func nextPermutation(nums []int) {
    k := -1
    j := len(nums) - 1
    for j > 0 {
        if nums[j-1] < nums[j] {
            k = j - 1
            break
        }
        j--
    }
    j = len(nums) - 1
    if k != -1 {
        for j > -1 {
            if nums[j] > nums[k] {
                nums[j], nums[k] = nums[k], nums[j]
                k++
                break
            }
            j--
        }
    } else {
        k = 0
    }
    j = len(nums) - 1
    for k < j {
        nums[k], nums[j] = nums[j], nums[k]
        k++
        j--
    }
}
