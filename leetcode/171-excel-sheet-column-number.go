// #string

func titleToNumber(columnTitle string) int {
    out := 0

    for _, letter := range columnTitle {
        out = (out * 26) + int(letter - 'A') + 1
    }

    return out
}
