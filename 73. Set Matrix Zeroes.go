package main

func setZeroes(matrix [][]int) {
    rowZeros := make(map[int]bool)
    colZeros := make(map[int]bool)
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                if !rowZeros[i] {
                    rowZeros[i] = true
                }
                if !colZeros[j] {
                    colZeros[j] = true
                }
            }
        }
    }
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            if rowZeros[i] || colZeros[j]{
                matrix[i][j] = 0
            }
        }
    }
}
