package main

import "strconv"

func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    str := countAndSay(n - 1)
    var result string
    slow := 0
    fast := 0
    for fast < len(str) {
        count := 0
        for fast < len(str) && str[fast] == str[slow] {
            count++
            fast++
        }
        result += strconv.Itoa(count) + string(str[slow])
        slow = fast
    }
    return result
}
