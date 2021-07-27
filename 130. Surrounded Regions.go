package main

func solve(board [][]byte) {
    for i := range board {
        for j := range board[i] {
            if i == 0 || j == 0 {
                infect(board, i, j)
            }
        }
    }
    for i := range board {
        for j := range board[i] {
            if board[i][j] == '1' {
                board[i][j] = 'O'
            } else {
                board[i][j] = 'X'
            }
        }
    }
}

func infect(board [][]byte, i, j int) {
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
        return
    }
    board[i][j] = '1'
    infect(board, i+1, j)
    infect(board, i-1, j)
    infect(board, i, j+1)
    infect(board, i, j-1)
}
