// #string

import (
    "strings"
)

func cellsInRange(s string) []string {
    out := []string{}
    tokens := strings.Split(s, ":")

    for col := tokens[0][0]; col <= tokens[1][0]; col++ {
        str := string(col)

        for row := tokens[0][1]; row <= tokens[1][1]; row++ {
            out = append(out, str + string(row))            
        }
    }

    return out
}
