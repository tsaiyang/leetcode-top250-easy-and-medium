package main

import (
    "strconv"
    "strings"
)

func compareVersion(version1 string, version2 string) int {
    v1 := strings.Split(version1, ".")
    v2 := strings.Split(version2, ".")
    index1, index2 := 0, 0
    n1, n2 := 0, 0
    for index1 < len(v1) || index2 < len(v2) {
        if index1 >= len(v1) {
            n1, _ = strconv.Atoi(v1[index1])
        } else {
            n1 = 0
        }

        if index2 >= len(v2) {
            n2, _ = strconv.Atoi(v2[index2])
        } else {
            n2 = 0
        }

        if n1 > n2 {
            return 1
        }
        if n1 < n2 {
            return -1
        }
        index1++
        index2++
    }
    return 0
}
