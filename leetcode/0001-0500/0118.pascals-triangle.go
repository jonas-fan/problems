// #array

func generate(numRows int) [][]int {
    out := make([][]int, numRows)

    for i := 0; i < numRows; i++ {
        columns := i + 1

        out[i] = make([]int, columns)

        for col := 0; col < columns; col++ {
            if col == 0 || (col + 1 == columns) {
                out[i][col] = 1
            } else {
                out[i][col] = out[i-1][col-1] + out[i-1][col]
            }
        }
    }

    return out
}
