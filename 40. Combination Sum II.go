package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var result [][]int
    arr := make([]int, 0)
    var fn func(level int)
    fn = func(level int) {
        if arraySum(arr) > target {
            return
        }
        if level == len(candidates) {
            if target == arraySum(arr) {
                tmp := make([]int, len(arr))
                copy(tmp, arr)
                result = append(result, tmp)
            }
            return
        }
        arr = append(arr, candidates[level])
        if arraySum(arr) <= target {
            fn(level + 1)
        }
        arr = arr[:len(arr)-1]

        index := level
        for index < len(candidates) {
            if index != level && candidates[index] != candidates[index-1] {
                break
            }
            index++
        }
        fn(index)
    }
    fn(0)
    return result
}

func arraySum(arr []int) int {
    sum := 0
    for i := range arr {
        sum += arr[i]
    }
    return sum
}
