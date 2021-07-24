package main

import "math"

func divide(dividend int, divisor int) int {
    if dividend == 0 {
        return 0
    }
    if divisor == 1 {
        return dividend
    }
    if divisor == -1 {
        if dividend > math.MinInt32 {
            return -dividend
        }
        return math.MaxInt32
    }

    a := dividend
    b := divisor
    same := true
    if a > 0 && b < 0 || a < 0 && b > 0 {
        same = false
    }
    if a < 0 {
        a = -a
    }
    if b < 0 {
        b = -1
    }
    result := div(int64(a), int64(b))
    if same && result > math.MaxInt32 {
        return math.MinInt32
    }
    if same {
        return result
    }
    return -result
}

func div(a, b int64) int {
    if a < b {
        return 0
    }
    count := 1
    tb := b
    for tb+tb < a {
        count = count + count
        tb = tb + tb
    }
    return count + div(a-tb, b)
}
