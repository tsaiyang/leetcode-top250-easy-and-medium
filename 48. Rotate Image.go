package main

func rotate(matrix [][]int) {
    startRow, startCol, endRow, endCol := 0, 0, len(matrix)-1, len(matrix[0])-1
    for startRow < endRow {
        rotateEdge(matrix, startRow, startCol, endRow, endCol)
        startRow++
        startCol++
        endRow--
        endCol--
    }
}

func rotateEdge(matrix [][]int, startRow, startCol, endRow, endCol int) {
    length := endRow - startRow
    for i := 0; i < length; i++ {
        tmp := matrix[startRow][startCol+i]
        matrix[startRow][startCol+i] = matrix[endRow-i][startCol]
        matrix[endRow-i][startCol] = matrix[endRow][endCol-i]
        matrix[endRow][endCol-i] = matrix[startRow+i][endCol]
        matrix[startRow+i][endCol] = tmp
    }
}
