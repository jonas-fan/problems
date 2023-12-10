// #string

import (
    "strconv"
    "strings"
)

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func sum(digits string) string {
    num := 0

    for _, digit := range digits {
        num += int(digit - '0')
    }

    return strconv.Itoa(num)
}

func digitSum(s string, k int) string {
    if len(s) <= k {
        return s
    }

    builder := &strings.Builder{}

    for i := 0; i < len(s); i += k {
        end := min(len(s), i+k)

        builder.WriteString(sum(s[i:end]))
    }

    return digitSum(builder.String(), k)
}
