// #string

import (
    "strings"
)

func prefixCount(words []string, pref string) int {
    out := 0

    for _, word := range words {
        if strings.HasPrefix(word, pref) {
            out++
        }
    }

    return out
}
