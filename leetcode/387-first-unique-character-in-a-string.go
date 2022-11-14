// #string

func counts(chars string) [256]int {
    out := [256]int{}

    for _, char := range chars {
        out[int(char)]++
    }

    return out
}

func firstUniqChar(s string) int {
    visits := counts(s)

    for i, char := range s {
        if visits[int(char)] == 1 {
            return i
        }
    }

    return -1
}
