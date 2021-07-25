package main

import "sort"

func merge(intervals [][]int) [][]int {
    sort.Sort(Matrix(intervals))
    result := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] > result[len(result)-1][1] {
            result = append(result, intervals[i])
        } else if intervals[i][1] > result[len(result)-1][1] {
            result[len(result)-1][1] = intervals[i][1]
        }
    }
    return result
}

type Matrix [][]int

func (m Matrix) Len() int {
    return len(m)
}

func (m Matrix) Less(i, j int) bool {
    if m[i][0] < m[j][0] {
        return true
    }
    if m[i][0] == m[j][0] && m[i][1] < m[j][1] {
        return true
    }
    return false
}

func (m Matrix) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}
