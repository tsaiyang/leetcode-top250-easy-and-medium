package main

func countPrimes(n int) int {
    primeList := make([]bool, n)
    for i := range primeList {
        primeList[i] = true
    }
    count := 0
    for i := 2; i < n; i++ {
        if primeList[i] {
            count++
            for j := i * 2; j < n; j += i {
                primeList[j] = false
            }
        }
    }
    return count
}
