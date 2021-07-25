package main

func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) < 1 {
        return [][]int{newInterval}
    }
    var result [][]int
    if newInterval[1] < intervals[0][0] {
        result = append(result, newInterval)
        result = append(result, intervals...)
        return result
    }
    for i := range intervals {
        if intervals[i][1] < newInterval[0] {
            result = append(result, intervals[i])
            continue
        }
        if intervals[i][0] <= newInterval[0] && intervals[i][1] >= newInterval[1] {
            return intervals
        }
        num1 := min(intervals[i][0], newInterval[0])
        for j := i; j < len(intervals); j++ {
                if newInterval[1] > intervals[j][1] {
                continue
            }
                if newInterval[1] < intervals[j][0] {
                result = append(result, intervals[:i]...)
                result = append(result, []int{num1, newInterval[1]})
                result = append(result, intervals[j:]...)
                return result
            }
                result = append(result, intervals[:i]...)
                result = append(result, []int{num1, intervals[j][1]})
            if j+1 < len(intervals) {
                result = append(result, intervals[j+1:]...)
            }
            return result
        }
    }
    return append(intervals, newInterval)
}
