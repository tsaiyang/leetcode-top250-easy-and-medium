package main

func isHappy(n int) bool {
    for i := 0; i <= 10000; i++ {
        var sum int
        for n > 0 {
            sum += (n % 10) * (n % 10)
            n /= 10
        }
        if sum == 1 {
            return true
        }
        n = sum
    }
    return false
}
