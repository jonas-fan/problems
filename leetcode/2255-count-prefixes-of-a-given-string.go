// #string

import (
    "strings"
)

func countPrefixes(words []string, s string) int {
    out := 0

    for _, word := range words {
        if strings.HasPrefix(s, word) {
            out++
        }
    }

    return out
}
