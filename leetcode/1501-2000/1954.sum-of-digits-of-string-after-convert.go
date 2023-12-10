// #string

import (
    "strconv"
)

func reverse(letters []rune) []rune {
    size := len(letters)

    for i := 0; i < size / 2; i++ {
        letters[i], letters[size-i-1] = letters[size-i-1], letters[i]
    }

    return letters
}

func convert(letters []rune) []rune {
    nums := make([]rune, 0, len(letters))

    for _, letter := range letters {
        code := int(letter - 'a') + 1

        if code > 9 {
            nums = append(nums, rune(code / 10) + '0')
            code %= 10
        }

        nums = append(nums, rune(code) + '0')
    }

    return nums
}

func transform(digits []rune, k int) int {
    if k < 1 {
        val, _ := strconv.Atoi(string(reverse(digits)))

        return val
    }

    out := make([]rune, 0, len(digits))
    num := 0

    for _, digit := range digits {
        num += int(digit - '0')
    }

    for num > 0 {
        out = append(out, rune(num % 10) + '0')
        num /= 10
    }

    return transform(out, k - 1)
}

func getLucky(s string, k int) int {
    return transform(convert([]rune(s)), k)
}
