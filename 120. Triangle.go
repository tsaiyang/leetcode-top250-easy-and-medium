package main

import "math"

func minimumTotal(triangle [][]int) int {
    for i := 1; i < len(triangle); i++ {
        if i == 1 {
            triangle[1][0] += triangle[0][0]
            triangle[1][1] += triangle[0][0]
            continue
        }
        for j := 0; j < i+1; j++ {
            if j == 0 {
                triangle[i][j] += triangle[i-1][j]
            } else if j == i {
                triangle[i][j] += triangle[i-1][j-1]
            } else {
                triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
            }
        }
    }
    result := math.MaxInt32
    for i := 0; i < len(triangle[len(triangle)-1]); i++ {
        result = min(result, triangle[len(triangle)-1][i])
    }
    return result
}
