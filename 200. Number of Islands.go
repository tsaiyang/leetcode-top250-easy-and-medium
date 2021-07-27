package main

func numIslands(grid [][]byte) int {
    var result int
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == '1' {
                result++
                infect1(grid, i, j)
            }
        }
    }
    return result
}

func infect1(grid [][]byte, i, j int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
        return
    }
    infect1(grid, i+1, j)
    infect1(grid, i-1, j)
    infect1(grid, i, j+1)
    infect1(grid, i, j-1)
}
