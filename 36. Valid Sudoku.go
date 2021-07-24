package main

func isValidSudoku(board [][]byte) bool {
    m := make(map[int]map[int]map[byte]bool)
    for i := range board {
        for j := range board[i] {
            if board[i][j] == ' ' {
                continue
            }
            if m[0] == nil {
                m[0] = make(map[int]map[byte]bool)
            }
            if m[0][i] == nil {
                m[0][i] = make(map[byte]bool)
            }
            if m[0][i][board[i][j]] {
                return false
            }
            m[0][i][board[i][j]] = true

            if m[1] == nil {
                m[1] = make(map[int]map[byte]bool)
            }
            if m[1][j] == nil {
                m[1][j] = make(map[byte]bool)
            }
            if m[1][j][board[i][j]] {
                return false
            }
            m[1][j][board[i][j]] = true

            chunk := (i/3)*3 + j/3
            if m[2] == nil {
                m[2] = make(map[int]map[byte]bool)
            }
            if m[2][chunk] == nil {
                m[2][chunk] = make(map[byte]bool)
            }
            if m[2][chunk][board[i][j]] {
                return false
            }
            m[2][chunk][board[i][j]] = true
        }
    }
    return true
}
