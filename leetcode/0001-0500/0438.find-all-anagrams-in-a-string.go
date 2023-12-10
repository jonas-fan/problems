// #string

func appears(chars string) [26]int {
    out := [26]int{}

    for _, char := range chars {
        out[char-'a']++
    }

    return out
}

func equals(lhs []int, rhs []int) bool {
    if len(lhs) != len(rhs) {
        return false
    }

    for i := range lhs {
        if lhs[i] != rhs[i] {
            return false
        }
    }

    return true
}

func findAnagrams(s string, p string) []int {
    if len(s) < len(p) {
        return []int{}
    }

    out := []int{}
    want := appears(p)
    seen := appears(s[:len(p)])

    if equals(seen[:], want[:]) {
        out = append(out, 0)
    }

    for i := 1; i <= len(s) - len(p); i++ {
        seen[s[i-1]-'a']--
        seen[s[i-1+len(p)]-'a']++

        if equals(seen[:], want[:]) {
            out = append(out, i)
        }
    }

    return out
}
