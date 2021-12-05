func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    return x == reverse(x)
}

func reverse(num int) int {
    out := 0

    for num > 0 {
        out *= 10
        out += num % 10
        num /= 10
    }

    return out
}
