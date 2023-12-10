// #string

func checkDistances(s string, distance []int) bool {
    for i, letter := range s {
        digit := int(letter - 'a')
        lindex := i - distance[digit] - 1
        rindex := i + distance[digit] + 1

        if lindex >= 0 && rune(s[lindex]) == letter {
            continue
        } else if rindex < len(s) && rune(s[rindex]) == letter {
            continue
        }

        return false
    }

    return true
}
