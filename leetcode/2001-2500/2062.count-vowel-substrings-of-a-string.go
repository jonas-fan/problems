// #string

func hashRune(r rune) uint {
    return uint(0x1) << (r - 'a')
}

func hashString(str string) uint {
    var out uint

    for _, char := range str {
        out |= hashRune(char)
    }

    return out
}

func countVowelSubstrings(word string) int {
    out := 0
    mask := hashString("aeiou")

    for i := 0; i < len(word); i++ {
        for j := i + 4; j < len(word); j++ {
            code := hashString(word[i:j+1])

            if code & ^mask > 0 {
                // not matched
            } else if code & mask == mask {
                out++
            }
        }
    }

    return out
}
