package main

func majorityElement(nums []int) int {
    times := 0
    result := 0
    for i, num := range nums {
        if times == 0 {
            result = num
            times = 1
        } else {
            times--
            i--
        }
    }
    return result
}
