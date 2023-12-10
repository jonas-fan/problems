// #string

import (
    "math"
)

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func commonChars(words []string) []string {
    if len(words) == 0 {
        return words
    }

    seen := make([][26]int, len(words))

    for i, word := range words {
        for _, letter := range word {
            seen[i][int(letter-'a')]++
        }
    }

    out := make([]string, 0, len(words[0]))

    for i := 0; i < 26; i++ {
        count := math.MaxInt

        for j := 0; j < len(seen) && count > 0; j++ {
            count = min(count, seen[j][i])
        }

        for ; count > 0; count-- {
            out = append(out, string(i+'a'))
        }
    }

    return out
}
