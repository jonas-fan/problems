// #string

func reverse(chars []rune) {
    size := len(chars)

    for i := 0; i < size / 2; i++ {
        chars[i], chars[size-i-1] = chars[size-i-1], chars[i]
    }
}

func reverseWordsRune(chars []rune) []rune {
    begin := 0

    for i := 0; i < len(chars); i++ {
        if chars[i] == ' ' {
            reverse(chars[begin:i])
            begin = i + 1
        }
    }

    reverse(chars[begin:])

    return chars
}

func reverseWords(s string) string {
    return string(reverseWordsRune([]rune(s)))
}
