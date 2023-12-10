// #math

func isUgly(n int) bool {
    for n >= 2 {
        if n & 0x1 == 1 {
            break
        }

        n = n >> 1
    }

    for n >= 5 {
        if n % 5 != 0 {
            break
        }

        n /= 5
    }

    for n >= 3 {
        if n % 3 != 0 {
            break
        }

        n /= 3
    }

    if n < 1 {
        return false
    } else if n > 1 {
        return false
    }

    return true
}
