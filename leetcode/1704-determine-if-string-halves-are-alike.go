// #string

func countVowels(str string) int {
    out := 0

    for _, letter := range str {
        if letter >= 'A' && letter <= 'Z' {
            letter = letter - 'A' + 'a'
        }

        switch letter {
        case 'a':
            fallthrough
        case 'e':
            fallthrough
        case 'i':
            fallthrough
        case 'o':
            fallthrough
        case 'u':
            out++
        }
    }

    return out
}

func halvesAreAlike(s string) bool {
    bound := len(s) >> 1
    lhs, rhs := s[:bound], s[bound:]

    return countVowels(lhs) == countVowels(rhs)
}
