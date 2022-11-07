// #string

import (
    "strconv"
)

func maximum69Number (num int) int {
    digits := []rune(strconv.Itoa(num))

    for i := 0; i < len(digits); i++ {
        if digits[i] == '6' {
            digits[i] = '9'
            break
        }
    }

    val, _ := strconv.Atoi(string(digits))

    return val
}
