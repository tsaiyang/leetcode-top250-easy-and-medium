package main

import "strings"

func reverseWords(s string) string {
    parts := strings.Split(s, " ")
    strs := make([]string, 0, len(parts))
    for i := range parts {
        if parts[i] != "" {
            strs = append(strs, parts[i])
        }
    }
    var result string
    for i := range strs {
        strs[i] = reverseStr(strs[i])
        result += strs[i] + " "
    }
    result = result[:len(result)-1]
    return reverseStr(result)
}

func reverseBytes(bytes []byte, i, j int) {
    for i < j {
        bytes[i], bytes[j] = bytes[j], bytes[i]
        i++
        j--
    }
}
