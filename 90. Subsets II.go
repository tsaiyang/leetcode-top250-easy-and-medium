package main

func subsetsWithDup(nums []int) [][]int {
    var result [][]int
    var arr []int
    var fn func(level int)
    fn = func(level int) {
        if level == len(nums) {
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            result = append(result, tmp)
            return
        }
        arr = append(arr, nums[level])
        fn(level + 1)
        arr = arr[:len(arr)-1]
        index := level
        for ; index < len(nums); index++ {
            if index != level && nums[index] != nums[index-1] {
                break
            }
        }
        fn(index)
    }
    fn(0)
    return result
}
