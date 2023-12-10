// #math

func hammingWeight(num uint32) int {
    out := 0

    for num > 0 {
        out += int(num & 0x1)
        num >>= 1
    }

    return out
}
