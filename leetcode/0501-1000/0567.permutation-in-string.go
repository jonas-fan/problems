// #string

func counts(chars string) [26]int {
    out := [26]int{}

    for _, char := range chars {
        out[char-'a']++
    }

    return out
}

func equals(lhs [26]int, rhs [26]int) bool {
    for i := range lhs {
        if lhs[i] != rhs[i] {
            return false
        }
    }

    return true
}

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    size := len(s1)
    want := counts(s1)
    seen := counts(s2[:size])

    for i := 0; i < len(s2) - size; i++ {
        if equals(seen, want) {
            return true
        }

        seen[s2[i]-'a']--
        seen[s2[i+size]-'a']++
    }

    return equals(seen, want)
}
