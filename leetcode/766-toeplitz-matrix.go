// #array

func isToeplitzMatrix(matrix [][]int) bool {
    for row := len(matrix) - 1; row > 0; row-- {
        for col := 1; col < len(matrix[row]); col++ {
            if matrix[row][col] != matrix[row-1][col-1] {
                return false
            }
        }
    }

    return true
}
