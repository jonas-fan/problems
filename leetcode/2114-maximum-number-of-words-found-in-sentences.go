// #string

import (
    "strings"
)

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func mostWordsFound(sentences []string) int {
    out := 0

    for _, sentence := range sentences {
        out = max(out, 1+strings.Count(sentence, " "))
    }

    return out
}
