// #string

func counts(chars string) [256]int {
    out := [256]int{}

    for _, char := range chars {
        out[char]++
    }

    return out
}

func canConstruct(ransomNote string, magazine string) bool {
    needs := counts(ransomNote)
    seen := counts(magazine)

    for char, count := range needs {
        if count == 0 {
            continue
        }

        if seen[char] < count {
            return false
        }
    }

    return true
}
