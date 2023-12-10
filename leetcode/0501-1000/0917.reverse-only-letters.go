// #string

func isalpha(letter rune) bool {
    if letter >= 'a' && letter <= 'z' {
        return true
    } else if letter >= 'A' && letter <= 'Z' {
        return true
    }

    return false
}

func reverse(letters []rune) []rune {
    left := 0
    right := len(letters) - 1

    for left < right {
        if !isalpha(letters[left]) {
            left++
            continue
        } else if !isalpha(letters[right]) {
            right--
            continue
        }

        letters[left], letters[right] = letters[right], letters[left]
        left++
        right--
    }

    return letters
}

func reverseOnlyLetters(s string) string {
    return string(reverse([]rune(s)))
}
