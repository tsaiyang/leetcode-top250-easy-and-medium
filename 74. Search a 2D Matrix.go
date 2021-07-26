package main

func searchMatrix(matrix [][]int, target int) bool {
    row := len(matrix)
    col := len(matrix[0])
    i := 0
    j := row*col - 1
    mid := 0
    for i+1 < j {
        mid = (i + j) / 2
        if target == matrix[mid/col][mid%col] {
            return true
        }
        if target > matrix[mid/col][mid%col] {
            i = mid + 1
        } else {
            j = mid - 1
        }
    }
    if matrix[i/col][i%col] == target {
        return true
    }
    if matrix[j/col][j%col] == target {
        return true
    }
    return false
}
