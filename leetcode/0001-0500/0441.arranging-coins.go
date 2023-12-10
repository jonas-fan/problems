// #math #binary-search

func arrangeCoins(n int) int {
    if n < 2 {
        return n
    }

    left := 1
    right := n

    n = n << 1

    for left < right {
        mid := left + (right - left) >> 1

        if mid + mid * mid == n {
            return mid
        } else if mid + mid * mid < n {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return left - 1
}
