// #string

func counts(letters []rune) [26]int {
    out := [26]int{}

    for _, letter := range letters {
        out[int(letter-'a')]++
    }

    return out
}

func findTheDifference(s string, t string) byte {
    lhs := counts([]rune(s))
    rhs := counts([]rune(t))

    for i := 0; i < len(lhs); i++ {
        if lhs[i] + 1 == rhs[i] {
            return byte(i + 'a')
        }
    }

    return byte(0)
}
