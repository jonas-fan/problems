// #string

func shrink(email string) string {
    out := make([]rune, 0, len(email))
    ignore := false

    for i, letter := range email {
        if letter == '@' {
            out = append(out, []rune(email[i:])...)
            break
        } else if letter == '.' {
            // ignore
        } else if letter == '+' {
            ignore = true
        } else if !ignore {
            out = append(out, letter)
        }
    }

    return string(out)
}

func numUniqueEmails(emails []string) int {
    seen := map[string]int{}

    for _, email := range emails {
        seen[shrink(email)]++
    }

    return len(seen)
}
