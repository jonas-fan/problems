// #string #stack

func removeDuplicates(s string) string {
    stack := make([]rune, 0, len(s))

    for _, letter := range s {
        size := len(stack)

        if size == 0 {
            // empty stack
        } else if stack[size-1] == letter {
            stack = stack[:size-1]
            continue
        }

        stack = append(stack, letter)
    }

    return string(stack)
}
