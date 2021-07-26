package main

func exist(board [][]byte, word string) bool {
    for i := range board {
        for j := range board[i] {
            occupied := make([][]bool, len(board))
            for i := range occupied {
                occupied[i] = make([]bool, len(board[i]))
            }

            var fn func(level, x, y int) bool
            fn = func(level, x, y int) bool {
                if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || occupied[x][y] || word[level] != board[x][y] {
                    return false
                }
                if level == len(word) {
                    return true
                }
                occupied[x][y] = true
                return fn(level+1, x+1, y) || fn(level+1, x, y) || fn(level+1, x, y+1) || fn(level+1, x, y-1)
            }
            if fn(0, i, j) {
                return true
            }
        }
    }
    return false
}
