// #string

func isOdd(num int) bool {
    return num & 0x1 == 1
}

func squareIsWhite(coordinates string) bool {
    col := int(coordinates[0]) - 'a'
    row := int(coordinates[1]) - '1'

    if isOdd(row) {
        return !isOdd(col)
    }

    return isOdd(col)
}
