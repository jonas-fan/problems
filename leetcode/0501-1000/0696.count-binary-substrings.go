// #string

func foo(digits []rune, first int , last int) int {
    out := 0

    if digits[first] == digits[last] {
        return out
    }

    for {
        if first--; first < 0 {
            break
        } else if digits[first] != digits[first+1] {
            break
        }

        if last++; last >= len(digits) {
            break
        } else if digits[last] != digits[last-1] {
            break
        }

        out++
    }

    return out + 1
}

func countBinarySubstrings(s string) int {
    out := 0
    digits := []rune(s)

    for i := 0; i < len(digits) - 1; i++ {
        if digits[i] != digits[i+1] {
            out += foo(digits, i, i+1)
        }
    }

    return out
}
