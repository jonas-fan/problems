// #bits

func reverseBits(num uint32) uint32 {
    var out uint32

    for i := 0; i < 32; i++ {
        out = out << 1
        out = out | (num & 0x1)
        num = num >> 1
    }

    return out
}
