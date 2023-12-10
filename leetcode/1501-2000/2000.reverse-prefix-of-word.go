// #string

func reverse(word []rune) []rune {
    for i := 0; i < len(word) / 2; i++ {
        word[i], word[len(word)-1-i] = word[len(word)-1-i], word[i]
    }

    return word
}

func reversePrefix(word string, ch byte) string {
    for i, letter := range word {
        if letter == rune(ch) {
            prefix := []rune(word[:i+1])

            return string(reverse(prefix)) + word[i+1:]
        }
    }

    return word
}
