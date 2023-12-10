// #array #graph

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func walk(grid [][]int, row int, col int, visited [][]bool) int {
    if row < 0 || col < 0 {
        return 0
    } else if row >= len(grid) || col >= len(grid[0]) {
        return 0
    } else if grid[row][col] == 0 {
        return 0
    } else if visited[row][col] {
        return 0
    }

    visited[row][col] = true

    out := 1
    out += walk(grid, row-1, col, visited)
    out += walk(grid, row+1, col, visited)
    out += walk(grid, row, col-1, visited)
    out += walk(grid, row, col+1, visited)

    return out
}

func maxAreaOfIsland(grid [][]int) int {
    visited := make([][]bool, len(grid))

    for i := range visited {
        visited[i] = make([]bool, len(grid[0]))
    }

    out := 0

    for row := 0; row < len(grid); row++ {
        for col := 0; col < len(grid[row]); col++ {
            if visited[row][col] {
                continue
            }

            out = max(out, walk(grid, row, col, visited))
        }
    }

    return out
}
