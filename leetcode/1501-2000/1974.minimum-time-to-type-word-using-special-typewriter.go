// #string

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func minTimeToType(word string) int {
    out := len(word)
    pointer := 'a'

    for _, letter := range word {
        lseconds := (int(letter-pointer) + 26) % 26
        rseconds := (int(pointer-letter) + 26) % 26

        out += min(lseconds, rseconds)
        pointer = letter
    }

    return out
}
