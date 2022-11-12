// #array

func matrixReshape(mat [][]int, r int, c int) [][]int {
    if len(mat) * len(mat[0]) != r * c {
        return mat
    }

    out := make([][]int, r)

    for i := 0; i < r; i++ {
        out[i] = make([]int, c)
    }

    index := 0

    for _, row := range mat {
        for _, val := range row {
            out[index/c][index%c] = val
            index++
        }
    }

    return out
}
