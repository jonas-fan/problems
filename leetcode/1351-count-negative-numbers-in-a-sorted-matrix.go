// #array

func countNegatives(grid [][]int) int {
    out := 0

    for row := 0; row < len(grid); row++ {
        for col := 0; col < len(grid[row]); col++ {
            if grid[row][col] < 0 {
                out += len(grid[row]) - col
                break
            }
        }
    }

    return out
}
