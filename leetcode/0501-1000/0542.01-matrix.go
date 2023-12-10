// #array

import (
    "math"
)

type Point struct {
    row int
    col int
}

func (p *Point) Up() *Point {
    return &Point{row: p.row-1, col: p.col}
}

func (p *Point) Down() *Point {
    return &Point{row: p.row+1, col: p.col}
}

func (p *Point) Left() *Point {
    return &Point{row: p.row, col: p.col-1}
}

func (p *Point) Right() *Point {
    return &Point{row: p.row, col: p.col+1}
}

func update(out [][]int, point *Point, moves int) bool {
    if point.row < 0 || point.row >= len(out) {
        return false
    } else if point.col < 0 || point.col >= len(out[0]) {
        return false
    } else if moves >= out[point.row][point.col] {
        return false
    }

    out[point.row][point.col] = moves

    return true
}

func updateMatrix(mat [][]int) [][]int {
    queue := make([]*Point, 0)
    out := make([][]int, len(mat))

    for row := range mat {
        out[row] = make([]int, len(mat[row]))

        for col := range mat[row] {
            if mat[row][col] > 0 {
                out[row][col] = math.MaxInt
            } else {
                queue = append(queue, &Point{row: row, col: col})
            }
        }
    }

    for len(queue) > 0 {
        next := make([]*Point, 0, len(queue) << 2)

        for _, point := range queue {
            moves := out[point.row][point.col] + 1

            if target := point.Up(); update(out, target, moves) {
                next = append(next, target)
            }

            if target := point.Down(); update(out, target, moves) {
                next = append(next, target)
            }

            if target := point.Left(); update(out, target, moves) {
                next = append(next, target)
            }

            if target := point.Right(); update(out, target, moves) {
                next = append(next, target)
            }
        }

        queue = next
    }

    return out
}
