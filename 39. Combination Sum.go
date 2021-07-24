package main

import "strconv"

func combinationSum(candidates []int, target int) [][]int {
    preSum := 0
    var result [][]int
    var arr []int
    var fn func(level int)
    fn = func(level int) {
        if level == len(candidates) {
            if target == preSum {
                tmp := make([]int, len(arr))
                copy(tmp, arr)
                result = append(result, tmp)
            }
            return
        }
        for i := 0; i*candidates[level]+preSum <= target; i++ {
            for j := 0; j < i; j++ {
                arr = append(arr, candidates[level])
            }
            preSum += i * candidates[level]
            fn(level + 1)
            for j := 0; j < i; j++ {
                arr = arr[:len(arr)-1]
            }
            preSum -= i * candidates[level]
        }
    }
    fn(0)
    return result
}

func array2String(arr []int) string {
    var result string
    for i := range arr {
        result += strconv.Itoa(arr[i]) + "-"
    }
    return result
}
