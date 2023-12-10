// #binary-search

func mySqrt(x int) int {
    left := 1
    right := x

    for left <= right {
        mid := left + (right - left) >> 1

        if mid * mid == x {
            return mid
        } else if mid * mid < x {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return left - 1
}
