package main

func getRow(rowIndex int) []int {
    if rowIndex < 1 {
        return []int{1}
    }
    if rowIndex < 2 {
        return []int{1, 1}
    }
    result := [][]int{{1}, {1, 1}}
    for i := 2; i < rowIndex + 1; i++ {
        arr := make([]int, i+1)
        for j := 0; j < i+1; j++ {
            if j == 0 || j == i {
                arr[j] = 1
            } else {
                arr[j] = result[i-1][j] + result[i-1][j-1]
            }
        }
        result = append(result, arr)
    }
    return result[len(result)-1]
}
