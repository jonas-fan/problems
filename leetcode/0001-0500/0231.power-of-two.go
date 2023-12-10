// #math

import (
    "math/bits"
)

func isPowerOfTwo(n int) bool {
    return bits.OnesCount(uint(n)) == 1
}
