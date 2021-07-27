package main

func rangeBitwiseAnd(left int, right int) int {
    var result int64
    for i := left; i <= right; i++ {
        result ^= int64(i)
    }
    return int(result)
}
