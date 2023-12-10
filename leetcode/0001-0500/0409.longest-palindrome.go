// #string

import (
    "unicode"
)

func longestPalindrome(s string) int {
    counts := [52]int{}

    for _, char := range s {
        if unicode.IsLower(char) {
            counts[int(char-'a')]++
        } else {
            counts[26+int(char-'A')]++
        }
    }

    out := 0

    for _, count := range counts {
        if count & 0x1 == 0 {
            out += count
        } else if out & 0x1 == 0 {
            out += count
        } else {
            out += count - 1
        }
    }

    return out
}
