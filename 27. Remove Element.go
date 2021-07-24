package main

func removeElement(nums []int, val int) int {
    i := 0
    j := len(nums) - 1
    for i < j {
        if nums[j] == val {
            j--
            continue
        }
        if nums[i] == val {
            nums[i],nums[j] = nums[j],nums[i]
        }
        i++
    }
    return i
}
