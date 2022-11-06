// #string

func max(lhs byte, rhs byte) byte {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
    maxTime := 0
    key := byte(0)

    for i := 0; i < len(releaseTimes); i++ {
        duration := releaseTimes[i]

        if i > 0 {
            duration -= releaseTimes[i-1]
        }

        if duration > maxTime {
            maxTime = duration
            key = keysPressed[i]
        } else if duration == maxTime {
            key = max(key, keysPressed[i])
        }
    }

    return key
}
