package main

func generateMatrix(n int) [][]int {
    startRow, startCol, endRow, endCol := 0, 0, n-1, n-1
    result := make([][]int, n)
    for i := range result {
        result[i] = make([]int, n)
    }
    count := 1
    for startRow <= endRow {
        if startRow == endRow {
            result[startRow][startCol] = count
            break
        }
        curRow := startRow
        curCol := startCol
        for curCol < endCol {
            result[startRow][curCol] = count
            count++
            curCol++
        }
        for curRow < endRow {
            result[curRow][endCol] = count
            count++
            curRow++
        }
        for curCol > startCol {
            result[endRow][curCol] = count
            count++
            curCol--
        }
        for curRow > startRow {
            result[curRow][startCol] = count
            count++
            curRow--
        }
        startRow++
        startCol++
        endRow--
        endCol--
    }
    return result
}
