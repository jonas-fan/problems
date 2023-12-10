// #string

func greatestLetter(s string) string {
    seen := make([]int, 52)

    for _, letter := range s {
        if letter >= 'a' && letter <= 'z' {
            seen[letter-'a']++
        } else {
            seen[letter-'A'+26]++
        }
    }

    for i := 25; i >= 0; i-- {
        if seen[i] == 0 {
            continue
        } else if seen[i+26] == 0 {
            continue
        }

        return string(i+'A')
    }

    return ""
}
