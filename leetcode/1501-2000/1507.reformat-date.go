// #string

import (
    "strings"
)

func toMonth(token string) int {
    switch strings.ToLower(token) {
    case "jan":
        return 1
    case "feb":
        return 2
    case "mar":
        return 3
    case "apr":
        return 4
    case "may":
        return 5
    case "jun":
        return 6
    case "jul":
        return 7
    case "aug":
        return 8
    case "sep":
        return 9
    case "oct":
        return 10
    case "nov":
        return 11
    case "dec":
        return 12
    default:
        return 0
    }
}

func toDay(token string) int {
    out := 0

    for _, letter := range token {
        if letter < '0' || letter > '9' {
            break
        }

        out = (out * 10) + int(letter - '0')
    }

    return out
}

func reformatDate(date string) string {
    tokens := strings.Split(date, " ")

    year := tokens[2]
    month := toMonth(tokens[1])
    day := toDay(tokens[0])

    return fmt.Sprintf("%s-%02d-%02d", year, month, day)
}
