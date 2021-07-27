package main

import "strings"

func isPalindrome1(s string) bool {
    s = strings.ToLower(s)
    i := 0
    j := len(s) - 1
    for i < j {
        if !isAlphaNumeric(s[i]) {
            i++
            continue
        }
        if !isAlphaNumeric(s[j]) {
            j--
            continue
        }
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func isAlphaNumeric(b byte) bool {
    if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
        return true
    }
    return false
}
