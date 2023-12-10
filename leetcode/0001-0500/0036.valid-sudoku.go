// #array

func validHorizontal(board [][]byte, row int, col int) bool {
    seen := [9]bool{}

    for ; col < 9; col++ {
        num := board[row][col]

        if num == '.' {
            continue
        }

        if seen[num-'1'] {
            return false
        }

        seen[num-'1'] = true
    }

    return true
}

func validVertical(board [][]byte, row int, col int) bool {
    seen := [9]bool{}

    for ; row < 9; row++ {
        num := board[row][col]

        if num == '.' {
            continue
        }

        if seen[num-'1'] {
            return false
        }

        seen[num-'1'] = true
    }

    return true
}

func validSquare(board [][]byte, row int, col int) bool {
    seen := [9]bool{}

    for i := row; i < row + 3; i++ {
        for j := col; j < col + 3; j++ {
            num := board[i][j]

            if num == '.' {
                continue
            }

            if seen[num-'1'] {
                return false
            }

            seen[num-'1'] = true
        }
    }

    return true
}

func isValidSudoku(board [][]byte) bool {
    for row := 0; row < len(board); row++ {
        if !validHorizontal(board, row, 0) {
            return false
        }
    }

    for col := 0; col < len(board[0]); col++ {
        if !validVertical(board, 0, col) {
            return false
        }
    }

    for row := 0; row < len(board); row += 3 {
        for col := 0; col < len(board[row]); col += 3 {
            if !validSquare(board, row, col) {
                return false
            }
        }
    }

    return true
}
