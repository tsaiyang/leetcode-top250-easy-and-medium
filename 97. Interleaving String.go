package main

func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) < 1 && len(s2) < 1 && len(s3) < 1 {
        return true
    }
    if len(s1)+len(s2) != len(s3) {
        return false
    }
    dp := make([][]bool, len(s1)+1)
    for i := range dp {
        dp[i] = make([]bool, len(s2)+1)
    }
    for i := 1; i < len(s1)+1; i++ {
        if s1[i-1] != s3[i-1] {
            break
        }
        dp[i][0] = true
    }
    for j := 1; j < len(s2)+1; j++ {
        if s2[j-1] != s3[j-1] {
            break
        }
        dp[0][j] = true
    }
    for i := 1; i < len(s1)+1; i++ {
        for j := 1; j < len(s2)+1; j++ {
            if (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1]) {
                dp[i][j] = true
            }
        }
    }
    return dp[len(s1)][len(s2)]
}
