// #math

func lastRemaining(n int) int {
    out := 1
    step := 1

    for n > 1 {
        // left to right
        {
            out = out + step
            n = n >> 1
            step = step << 1
        }

        if n <= 1 {
            return out
        }

        // right to left
        {
            if n & 0x1 == 1 {
                out = out + step
            }

            n = n >> 1
            step = step << 1
        }
    }

    return out
}
