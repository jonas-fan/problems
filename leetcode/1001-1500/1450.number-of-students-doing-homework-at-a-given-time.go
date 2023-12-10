// #array

func busyStudent(startTime []int, endTime []int, queryTime int) int {
    out := 0

    for i := 0; i < len(startTime); i++ {
        if startTime[i] > queryTime {
            continue
        } else if endTime[i] < queryTime {
            continue
        }

        out++
    }

    return out
}
