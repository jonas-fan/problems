func numberOfSteps(num int) int {
    count := 0

    for num > 0 {
        if num&0x1 == 0 {
            num >>= 1
        } else {
            num--
        }

        count++
    }

    return count
}
