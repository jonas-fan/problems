// #array

func makeSquare(size int) [][]int {
    out := make([][]int, 0, size)

    for i := 0; i < size; i++ {
        out = append(out, make([]int, size))
    }

    return out
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func maxMatrix(grid [][]int, row int, col int, size int) int {
    out := 0

    for i := row; i < row + size; i++ {
        for j := col; j < col + size; j++ {
            out = max(out, grid[i][j])
        }
    }

    return out
}

func largestLocal(grid [][]int) [][]int {
    maxLocal := makeSquare(len(grid) - 2)

    for row := 0; row < len(maxLocal); row++ {
        for col := 0; col < len(maxLocal); col++ {
            maxLocal[row][col] = maxMatrix(grid, row, col, 3)
        }
    }

    return maxLocal
}
