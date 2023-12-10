// #string

import (
    "unicode"
)

func isBad(lhs rune, rhs rune) bool {
    if unicode.IsUpper(lhs) {
        return false
    } else if unicode.IsLower(rhs) {
        return false
    } else if unicode.ToUpper(lhs) != rhs {
        return false
    }

    return true
}

func makeGoodRune(chars []rune) []rune {
    out := make([]rune, 0, len(chars))

    for i := 0; i < len(chars); i++ {
        if i + 1 == len(chars) {
            // last
        } else if isBad(chars[i], chars[i+1]) || isBad(chars[i+1], chars[i]) {
            i++
            continue
        }

        out = append(out, chars[i])
    }

    if len(out) < len(chars) {
        return makeGoodRune(out)
    }

    return out
}

func makeGood(s string) string {
    return string(makeGoodRune([]rune(s)))
}
