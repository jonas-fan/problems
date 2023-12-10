// #string

import (
    "strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
    out := 0

    for _, word := range strings.Split(text, " ") {
        if !strings.ContainsAny(word, brokenLetters) {
            out++
        }
    }

    return out
}
