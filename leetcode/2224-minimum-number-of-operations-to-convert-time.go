// #string

func abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}

func toMinutes(hhMM string) int {
    hours := int(hhMM[0] - '0') * 10 + int(hhMM[1] - '0')
    mins := int(hhMM[3] - '0') * 10 + int(hhMM[4] - '0')

    return hours * 60 + mins
}

func convertTime(current string, correct string) int {
    mins := abs(toMinutes(correct) - toMinutes(current))
    out := 0

    for mins > 0 {
        if mins >= 60 {
            out += mins / 60
            mins %= 60
        } else if mins >= 15 {
            out += mins / 15
            mins %= 15
        } else if mins >= 5 {
            out += mins / 5
            mins %= 5
        } else if mins >= 0 {
            out += mins
            mins = 0
        }
    }

    return out
}
