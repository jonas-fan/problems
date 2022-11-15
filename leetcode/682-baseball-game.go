// #array #stack

import (
    "strconv"
)

func calPoints(operations []string) int {
    stack := []int{}

    for _, ops := range operations {
        switch ops {
        case "+":
            top1 := stack[len(stack)-1]
            top2 := stack[len(stack)-2]
            stack = append(stack, top1 + top2)
        case "D":
            top := stack[len(stack)-1]
            stack = append(stack, top << 1)
        case "C":
            stack = stack[:len(stack)-1]
        default:
            val, _ := strconv.Atoi(ops)
            stack = append(stack, val)
        }
    }

    sum := 0

    for _, num := range stack {
        sum += num
    }

    return sum
}
