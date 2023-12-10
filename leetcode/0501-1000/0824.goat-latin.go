// #string

import (
    "strings"
)

func lower(r rune) rune {
    if r >= 'A' && r <= 'Z' {
        return r - 'A' + 'a'
    }

    return r
}

func isVowel(r rune) bool {
    switch lower(r) {
    case 'a', 'e', 'i', 'o', 'u':
        return true
    }

    return false
}

func toGoatLatin(sentence string) string {
    builder := &strings.Builder{}

    for i, word := range strings.Split(sentence, " ") {
        if i != 0 {
            builder.WriteRune(' ')            
        }

        leading := rune(word[0])

        if isVowel(leading) {
            builder.WriteString(word)
        } else {
            builder.WriteString(word[1:])
            builder.WriteRune(leading)
        }

        builder.WriteString("ma")
        builder.WriteString(strings.Repeat("a", i+1))
    }

    return builder.String()
}
