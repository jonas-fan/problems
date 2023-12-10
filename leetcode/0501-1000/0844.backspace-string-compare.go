// #string #stack

func canonical(input string) string {
    stack := make([]rune, 0, len(input))

    for _, char := range input {
        if char == '#' {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        } else {
            stack = append(stack, char)
        }
    }

    return string(stack)
}

func backspaceCompare(s string, t string) bool {
    return canonical(s) == canonical(t)
}
