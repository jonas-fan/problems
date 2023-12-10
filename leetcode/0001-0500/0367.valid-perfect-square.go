// #binary-search

func isPerfectSquare(num int) bool {
    left := 1
    right := num

    for left <= right {
        mid := left + (right - left) >> 1

        if mid * mid == num {
            return true
        } else if mid * mid < num {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return false
}
