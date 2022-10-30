// #string

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func areOccurrencesEqual(s string) bool {
    freq := 0
    counts := make([]int, 26)

    for _, letter := range s {
        index := letter - 'a'

        counts[index]++

        freq = max(freq, counts[index])
    }

    for _, count := range counts {
        if count == 0 {
            continue
        } else if count != freq {
            return false
        }
    }

    return true
}
