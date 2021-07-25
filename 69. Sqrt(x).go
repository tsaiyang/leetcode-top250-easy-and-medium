package main

func mySqrt(x int) int {
    i := 0
    j := 463401
    for i+1 < j {
        mid := (i + j) / 2
        if mid*mid == x {
            return mid
        }
        if mid*mid < x {
            i = mid
        } else {
            j = mid
        }
    }
    if j*j < x {
        return j
    }
    if i*i < x {
        return i
    }
    return i - 1
}
