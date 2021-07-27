package main

import (
    "sort"
    "strconv"
)

func largestNumber(nums []int) string {
    sort.Sort(NumSlice(nums))
    var result string
    for _, num := range nums {
        result += strconv.Itoa(num)
    }
    if num, _ := strconv.Atoi(result); num == 0 {
        return "0"
    }
    return result
}

type NumSlice []int

func (ns NumSlice) Len() int {
    return len(ns)
}

func (ns NumSlice) Less(i, j int) bool {
    str1 := strconv.Itoa(ns[i])
    str2 := strconv.Itoa(ns[j])
    return str1+str2 > str2+str1
}

func (ns NumSlice) Swap(i, j int) {
    ns[i], ns[j] = ns[j], ns[i]
}
