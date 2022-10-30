// #string

func digitCount(num string) bool {
    seen := make([]int, 10)

    for _, letter := range num {
        seen[letter-'0']++
    }

    for i, letter := range num {
        if seen[i] != int(letter-'0') {
            return false
        }
    }

    return true
}
