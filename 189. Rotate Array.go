package main

func rotate1(nums []int, k int) {
    if len(nums) < 2 || k < 0 {
        return
    }
    k = k % len(nums)
    reverseNums(nums, len(nums)-k, len(nums)-1)
    reverseNums(nums, 0, len(nums)-k-1)
    reverseNums(nums, 0, len(nums)-1)
}

func reverseNums(nums []int, i, j int) {
    for i < j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}
