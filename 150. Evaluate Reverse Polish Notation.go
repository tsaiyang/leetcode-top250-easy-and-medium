package main

import "strconv"

func evalRPN(tokens []string) int {
    stack := make([]int, len(tokens))
    for i := range tokens {
        if tokens[i] != "*" && tokens[i] != "/" && tokens[i] != "+" && tokens[i] != "-" {
            num, _ := strconv.Atoi(tokens[i])
            stack = append(stack, num)
        } else {
            num1 := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            num2 := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if tokens[i] == "+" {
                num1 += num2
            } else if tokens[i] == "*" {
                num1 *= num2
            } else if tokens[i] == "-" {
                num1 -= num2
            } else {
                num1 /= num2
            }
            stack = append(stack, num1)
        }
    }
    return stack[len(stack)-1]
}
