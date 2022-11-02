// #string

import (
    "strings"
)

func filter(number string) []rune {
    out := make([]rune, 0, len(number))

    for _, digit := range number {
        if digit == '-' || digit == ' ' {
            continue
        }

        out = append(out, digit)
    }

    return out
}

func reformatNumber(number string) string {
    digits := filter(number)
    builder := &strings.Builder{}
    
    for len(digits) > 0 {
        if len(digits) > 4 {
            builder.WriteString(string(digits[0:3]))
            builder.WriteRune('-')
            digits = digits[3:]
        } else if len(digits) == 4 {
            builder.WriteString(string(digits[0:2]))
            builder.WriteRune('-')
            digits = digits[2:]
        } else {
            builder.WriteString(string(digits))
            digits = nil
        }
    }

    return builder.String()
}
