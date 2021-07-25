package main

func spiralOrder(matrix [][]int) []int {
    startRow, startCol, endRow, endCol := 0, 0, len(matrix)-1, len(matrix[0])-1
    var result []int
    for startRow <= endCol && startCol <= endCol {
        result = append(result, matrixEdge(matrix, startRow, startCol, endRow, endCol)...)
        startRow++
        startCol++
        endRow--
        endCol--
    }
    return result
}

func matrixEdge(matrix [][]int, startRow, startCol, endRow, endCol int) []int {
    var result []int
    if startRow == endRow {
        for i := startCol; i <= endCol; i++ {
            result = append(result, matrix[startRow][i])
        }
        return result
    }
    if startCol == endCol {
        for i := startRow; i <= endRow; i++ {
            result = append(result, matrix[i][startCol])
        }
        return result
    }
    curRow := startRow
    curCol := startCol
    for curCol < endCol {
       result = append(result, matrix[startRow][curCol])
       curCol++
    }
    for curRow < endRow {
       result = append(result, matrix[curRow][endCol])
       curRow++
    }
    for curCol > startCol {
       result = append(result, matrix[endRow][curCol])
       curCol--
    }
    for curRow > startRow {
       result = append(result, matrix[curRow][startCol])
       curRow--
    }
    return result
}
