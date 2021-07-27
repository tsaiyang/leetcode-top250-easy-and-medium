package main

import "math"

func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    lowPrice := prices[0]
    result := math.MaxInt32
    for i := 1; i < len(prices); i++ {
        if prices[i] > lowPrice {
            result = min(result, prices[i]-lowPrice)
        } else {
            lowPrice = prices[i]
        }
    }
    return result
}
