// #string

import (
    "strconv"
    "unicode"
)

func findLast(chars []rune, r rune) int {
    for i := len(chars) - 1; i >= 0; i-- {
        if chars[i] == r {
            return i
        }
    }

    return -1
}

func decodeString(s string) string {
    stack := []rune{}

    for _, char := range s {
        if char == ']' {
            beginBracket := findLast(stack, '[')
            beginNum := beginBracket

            for beginNum > 0 {
                beginNum--

                if !unicode.IsDigit(stack[beginNum]) {
                    beginNum++
                    break
                }
            }

            num := string(stack[beginNum:beginBracket])
            token := string(stack[beginBracket+1:])

            stack = stack[:beginNum]

            if val, err := strconv.Atoi(num); err == nil {
                stack = append(stack, []rune(strings.Repeat(token, val))...)
            }
        } else {
            stack = append(stack, char)
        }
    }

    return string(stack)
}
