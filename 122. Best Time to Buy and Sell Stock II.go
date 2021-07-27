package main

func maxProfit1(prices []int) int {
    if len(prices) < 1 {
        return 0
    }
    onHand := -prices[0]
    offHand := 0
    for i := 1; i < len(prices); i++ {
        tmp := offHand
        offHand = max(offHand, prices[i]+onHand)
        onHand = max(onHand, tmp-prices[i])
    }
    return offHand
}
