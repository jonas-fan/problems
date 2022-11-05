// #string

import (
    "strings"
)

func tolower(r rune) rune {
    if r < 'A' || r > 'Z' {
        return r
    }

    return r - 'A' + 'a'
}

func toupper(r rune) rune {
    if r < 'a' || r > 'z' {
        return r
    }

    return r - 'a' + 'A'
}

func capitalize(word string) string {
    out := make([]rune, 0, len(word))

    for i, char := range word {
        if i == 0 {
            out = append(out, toupper(char))
        } else {
            out = append(out, tolower(char))
        }
    }

    return string(out)
}

func capitalizeTitle(title string) string {
    builder := &strings.Builder{}
    words := strings.Split(title, " ")

    for i, word := range words {
        if i > 0 {
            builder.WriteRune(' ')
        }

        if len(word) < 3 {
            builder.WriteString(strings.ToLower(word))
        } else {
            builder.WriteString(capitalize(word))
        }
    }

    return builder.String()
}
