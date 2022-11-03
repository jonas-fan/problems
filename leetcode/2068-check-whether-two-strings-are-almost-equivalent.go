/// #string

func abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}

func counts(word string) [26]int {
    out := [26]int{}

    for _, letter := range word {
        out[int(letter-'a')]++
    }

    return out
}

func checkAlmostEquivalent(word1 string, word2 string) bool {
    lhs := counts(word1)
    rhs := counts(word2)

    for i := range lhs {
        if abs(lhs[i] - rhs[i]) > 3 {
            return false
        }
    }

    return true
}
