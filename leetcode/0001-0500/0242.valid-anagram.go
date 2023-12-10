// #string

func counts(chars string) [256]int {
    out := [256]int{}

    for _, char := range chars {
        out[char]++
    }

    return out
}

func equals(lhs, rhs [256]int) bool {
    for i := 0; i < len(lhs); i++ {
        if lhs[i] != rhs[i] {
            return false
        }
    }

    return true
}

func isAnagram(s string, t string) bool {
    return equals(counts(s), counts(t))
}
