// #depth-first-search

import (
    "unicode"
)

func walk(str string, index int, chars []rune, out []string) []string {
    if index >= len(str) {
        return append(out, string(chars))
    }

    char := rune(str[index])

    if unicode.IsLetter(char) {
        out = walk(str, index+1, append(chars, unicode.ToLower(char)), out)
        out = walk(str, index+1, append(chars, unicode.ToUpper(char)), out)
    } else {
        out = walk(str, index+1, append(chars, char), out)
    }

    return out
}

func letterCasePermutation(s string) []string {
    return walk(s, 0, []rune{}, []string{})
}
