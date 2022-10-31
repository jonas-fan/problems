// #string

func rowNum(letter rune) int {
    if letter >= 'A' && letter <= 'Z' {
        letter = letter - 'A' + 'a'
    }

    switch letter {
    case 'q': fallthrough
    case 'w': fallthrough
    case 'e': fallthrough
    case 'r': fallthrough
    case 't': fallthrough
    case 'y': fallthrough
    case 'u': fallthrough
    case 'i': fallthrough
    case 'o': fallthrough
    case 'p': return 0

    case 'a': fallthrough
    case 's': fallthrough
    case 'd': fallthrough
    case 'f': fallthrough
    case 'g': fallthrough
    case 'h': fallthrough
    case 'j': fallthrough
    case 'k': fallthrough
    case 'l': return 1

    case 'z': fallthrough
    case 'x': fallthrough
    case 'c': fallthrough
    case 'v': fallthrough
    case 'b': fallthrough
    case 'n': fallthrough
    case 'm': return 2
    }

    return -1
}

func inOneRow(word string) bool {
    if len(word) == 0 {
        return true
    }

    row := rowNum(rune(word[0]))

    for _, letter := range word {
        if rowNum(letter) != row {
            return false
        }
    }

    return true
}

func findWords(words []string) []string {
    out := make([]string, 0, len(words))

    for _, word := range words {
        if inOneRow(word) {
            out = append(out, word)
        }
    }

    return out
}

