package main

import "strconv"

func restoreIpAddresses(s string) []string {
    var result []string
    for i := 0; i < len(s)-2; i++ {
        s1 := s[0:i]
        if !isValidIp(s1) {
            continue
        }
        for j := i + 1; j < len(s)-1; j++ {
            s2 := s[i:j]
            if !isValidIp(s2) {
                continue
            }
            for k := j + 1; k < len(s); k++ {
                s3 := s[j:k]
                if !isValidIp(s3) {
                    continue
                }
                s4 := s[k:]
                if !isValidIp(s4) {
                    continue
                }
                result = append(result, s1+"."+s2+"."+s3+"."+s4)
            }
        }
    }
    return result
}

func isValidIp(s string) bool {
    num, _ := strconv.Atoi(s)
    if len(s) > 3 || len(s) < 1 || num > 255 || (s[0] == '0' && len(s) > 1) {
        return false
    }
    return true
}
