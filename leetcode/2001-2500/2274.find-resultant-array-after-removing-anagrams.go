// #string

func counts(word string) [26]int {
    out := [26]int{}

    for _, char := range word {
        out[int(char-'a')]++
    }

    return out
}

func same(lhs [26]int, rhs [26]int) bool {
    for i := 0; i < len(lhs); i++ {
        if lhs[i] != rhs[i] {
            return false
        }
    }

    return true
}

func removeAnagrams(words []string) []string {
    out := make([]string, 0, len(words))
    lastSeen := counts("")

    for i := 0; i < len(words); i++ {
        seen := counts(words[i])

        if !same(seen, lastSeen) {
            lastSeen = seen
            out = append(out, words[i])
        }
    }

    return out
}
