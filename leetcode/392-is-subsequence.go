// #string

func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
        return true
    } else if len(t) == 0 {
        return false
    }

    index := 0

    for _, char := range t {
        if char != rune(s[index]) {
            continue
        }

        index++

        if index == len(s) {
            return true
        }
    }

    return false
}
