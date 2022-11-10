// #array

import (
    "sort"
)

func abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}

func minMovesToSeat(seats []int, students []int) int {
    sort.Ints(seats)
    sort.Ints(students)

    out := 0

    for i := 0; i < len(seats); i++ {
        out += abs(seats[i] - students[i])
    }

    return out
}
