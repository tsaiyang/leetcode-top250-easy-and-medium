package main

func generate(numRows int) [][]int {
    if numRows < 1 {
        return nil
    }
    if numRows < 2 {
        return [][]int{{1}}
    }
    if numRows < 3 {
        return [][]int{{1}, {1, 1}}
    }
    result := [][]int{{1}, {1, 1}}
    for i := 3; i < numRows+1; i++ {
        arr := make([]int, i)
        for j := 0; j < i; j++ {
            if j == 0 || j == i-1 {
                arr[j] = 1
            } else {
                arr[j] = result[i-2][j] + result[i-2][j-1]
            }
        }
        result = append(result, arr)
    }
    return result
}
