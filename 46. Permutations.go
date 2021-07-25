package main

func permute(nums []int) [][]int {
    var result [][]int
    var fn func(level int)
    fn = func(level int) {
        if level == len(nums) {
            tmp := make([]int, len(nums))
            copy(tmp, nums)
            result = append(result, tmp)
            return
        }
        for i := level; i < len(nums); i++ {
            nums[i], nums[level] = nums[level], nums[i]
            fn(level + 1)
            nums[i], nums[level] = nums[level], nums[i]
        }
    }
    fn(0)
    return result
}
