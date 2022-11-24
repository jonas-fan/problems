// #graph #depth-first-search

func walk(board [][]byte, row int, col int, word string) bool {
    if row < 0 || row >= len(board) {
        return false
    } else if col < 0 || col >= len(board[0]) {
        return false
    } else if board[row][col] == 0 {
        return false
    } else if board[row][col] != word[0] {
        return false
    }

    word = word[1:]

    if len(word) == 0 {
        return true
    }

    original := board[row][col]

    board[row][col] = 0

    if walk(board, row-1, col, word) {
        return true
    } else if walk(board, row+1, col, word) {
        return true
    } else if walk(board, row, col-1, word) {
        return true
    } else if walk(board, row, col+1, word) {
        return true
    }

    board[row][col] = original

    return false
}

func exist(board [][]byte, word string) bool {
    for row := 0; row < len(board); row++ {
        for col := 0; col < len(board[row]); col++ {
            if walk(board, row, col, word) {
                return true
            }
        }
    }

    return false
}
