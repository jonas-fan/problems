// #dynamic-programming

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func minimumTotal(triangle [][]int) int {
    sum := triangle

    for row := len(sum) - 2; row >= 0; row-- {
        for col := 0; col < len(sum[row]); col++ {
            sum[row][col] = min(sum[row][col] + sum[row+1][col],
                                sum[row][col] + sum[row+1][col+1])
        }
    }

    return sum[0][0]
}
