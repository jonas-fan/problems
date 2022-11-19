// #dynamic-programming

func uniquePaths(m int, n int) int {
    pathes := make([][]int, m)

    for i := range pathes {
        pathes[i] = make([]int, n)
    }

    for row := 0; row < m; row++ {
        pathes[row][0] = 1
    }

    for col := 0; col < n; col++ {
        pathes[0][col] = 1
    }

    for row := 1; row < m; row++ {
        for col := 1; col < n; col++ {
            pathes[row][col] = pathes[row-1][col] + pathes[row][col-1]
        }
    }

    return pathes[m-1][n-1]
}
