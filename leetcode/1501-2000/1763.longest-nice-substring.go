// #string

func isNice(letters []rune) bool {
    if len(letters) == 0 {
        return false
    }

    mask := [52]bool{}

    for _, letter := range letters {
        if letter >= 'a' && letter <= 'z' {
            mask[int(letter-'a')] = true
        } else {
            mask[int(letter-'A')+26] = true
        }
    }

    for i := 0; i < len(mask); i++ {
        if !mask[i] {
            continue
        }

        j := i

        if j < 26 {
            j += 26
        } else {
            j -= 26
        }

        if !mask[j] {
            return false
        }
    }

    return true
}

func longestNiceSubstring(s string) string {
    out := ""
    letters := []rune(s)

    for i := 0; i < len(letters); i++ {
        for j := len(letters); j > i + 1; j-- {
            subletters := letters[i:j]

            if isNice(subletters) {
                if len(subletters) > len(out) {
                    out = string(subletters)
                }
            }
        }
    }

    return out
}
