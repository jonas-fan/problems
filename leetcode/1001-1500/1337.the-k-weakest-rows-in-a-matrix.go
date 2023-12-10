// #array

import (
    "sort"
)

type Row struct {
    id      int
    solders int
}

func sum(nums []int) int {
    out := 0

    for _, num := range nums {
        out += num
    }

    return out
}

func kWeakestRows(mat [][]int, k int) []int {
    rows := make([]*Row, 0, len(mat))

    for i, row := range mat {
        rows = append(rows, &Row{
            id:      i,
            solders: sum(row),
        })
    }

    sort.Slice(rows, func(lhs int, rhs int) bool {
        if rows[lhs].solders == rows[rhs].solders {
            return rows[lhs].id < rows[rhs].id
        }

        return rows[lhs].solders < rows[rhs].solders
    })

    out := make([]int, 0, k)

    for i := 0; i < k; i++ {
        out = append(out, rows[i].id)
    }

    return out
}
