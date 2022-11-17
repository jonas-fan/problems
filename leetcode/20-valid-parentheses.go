// #stack

func isValid(s string) bool {
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    stack := []rune{}

    for _, char := range s {
        switch char {
        case ')':
            fallthrough
        case '}':
            fallthrough
        case ']':
            size := len(stack)

            if size == 0 {
                return false
            } else if pairs[char] != stack[size-1] {
                return false
            }

            stack = stack[:size-1]
        default:
            stack = append(stack, char)
        }
    }

    return len(stack) == 0
}
