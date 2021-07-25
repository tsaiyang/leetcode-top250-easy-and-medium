package main

func myPow(x float64, n int) float64 {
    if n < 0 && isZero(x) {
        return 0
    }
    isPositive := true
    if n < 0 {
        n = -n
        isPositive = false
    }
    num := pow(x, n)
    if !isPositive {
        return 1 / num
    }
    return num
}

func pow(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return x
    }
    num := pow(x, n/2)
    result := num * num
    if n%2 != 0 {
        result *= x
    }
    return result
}

func isZero(x float64) bool {
    if x > -0.000001 && x < 0.000001 {
        return true
    }
    return false
}
