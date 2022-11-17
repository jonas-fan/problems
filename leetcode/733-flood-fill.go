// #array #graph

func walk(image [][]int, row int, col int, want int, color int) [][]int {
    if want == color {
        return image
    } else if row < 0 || col < 0 {
        return image
    } else if row >= len(image) || col >= len(image[0]) {
        return image
    } else if image[row][col] != want {
        return image
    }

    image[row][col] = color

    image = walk(image, row-1, col, want, color)
    image = walk(image, row+1, col, want, color)
    image = walk(image, row, col-1, want, color)
    image = walk(image, row, col+1, want, color)

    return image
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    return walk(image, sr, sc, image[sr][sc], color)
}
