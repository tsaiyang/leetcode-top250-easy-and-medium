package main

func removeDuplicates2(nums []int) int {
    if len(nums) < 3 {
        return len(nums)
    }
    i := 2
    j := 2
    for j < len(nums) {
        if nums[j] != nums[i-2] {
            nums[i] = nums[j]
            i++
        }
        j++
    }
    return i + 1
}
