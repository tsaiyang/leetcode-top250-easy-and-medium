package main

func numTrees(n int) int {
    dp := make([]int, n+1)
    dp[0] = 1
    for i := 1; i < n+1; i++ {
        for j := i; j >= 0; j-- {
            dp[i] += dp[j-1] * dp[i-j]
        }
    }
    return dp[n]
}
