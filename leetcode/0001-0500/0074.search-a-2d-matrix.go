// #array

func searchMatrix(matrix [][]int, target int) bool {
    for i := len(matrix) - 1; i >= 0; i-- {
        if matrix[i][0] > target {
            continue
        }

        for _, num := range matrix[i] {
            if num == target {
                return true
            }
        }
    }

    return false
}
