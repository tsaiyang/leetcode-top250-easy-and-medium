package main

import "math"

func longestCommonPrefix(strs []string) string {
    minLen := math.MaxInt32
    for i := range strs {
        minLen = min(minLen, len(strs[i]))
    }
    result := make([]byte, 0, minLen)
    for i := 0; i < minLen; i++ {
        for j := range strs {
            if j != 0 && strs[j][i] != strs[j-1][i] {
                return string(result)
            }
        }
        result = append(result, strs[0][i])
    }
    return string(result)
}
