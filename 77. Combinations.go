package main

func combine(n int, k int) [][]int {
    var result [][]int
    var arr []int
    var fn func(level int)
    fn = func(level int) {
        if len(arr) == k {
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            result = append(result, tmp)
            return
        }
        if level == n+1 {
            return
        }
        arr = append(arr, level)
        fn(level + 1)
        arr = arr[:len(arr)-1]
        fn(level + 1)
    }
    fn(1)
    return result
}
