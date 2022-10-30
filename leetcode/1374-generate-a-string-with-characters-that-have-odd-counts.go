// #string

import (
    "strings"
)

func generateTheString(n int) string {
    if n & 0x1 == 0 {
        return strings.Repeat("a", n-1) + "b"
    }

    return strings.Repeat("a", n)
}
