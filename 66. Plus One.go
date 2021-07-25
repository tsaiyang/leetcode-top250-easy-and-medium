package main

func plusOne(digits []int) []int {
    flag := 1
    for i := len(digits) - 1;i >= 0;i-- {
        if flag != 1 {
            break
        }
        num := digits[i] + flag
        digits[i] = num % 10
        flag = num / 10
    }
    if flag == 1 {
        digits = append([]int{1}, digits...)
    }
    return digits
}
