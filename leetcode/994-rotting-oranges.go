// #array

const (
    FRESH = iota + 1
    ROTTEN
)

type Position struct {
    row int
    col int
}

func (p *Position) Up() *Position {
    return &Position{row: p.row-1, col: p.col}
}

func (p *Position) Down() *Position {
    return &Position{row: p.row+1, col: p.col}
}

func (p *Position) Left() *Position {
    return &Position{row: p.row, col: p.col-1}
}

func (p *Position) Right() *Position {
    return &Position{row: p.row, col: p.col+1}
}

func rot(grid [][]int, pos *Position) bool {
    if pos.row < 0 || pos.row >= len(grid) {
        return false
    } else if pos.col < 0 || pos.col >= len(grid[0]) {
        return false
    } else if grid[pos.row][pos.col] != FRESH {
        return false
    }

    grid[pos.row][pos.col] = ROTTEN

    return true
}

func orangesRotting(grid [][]int) int {
    fresh := 0
    queue := []*Position{}

    for row := range grid {
        for col := range grid[row] {
            if grid[row][col] == FRESH {
                fresh++
            } else if grid[row][col] == ROTTEN {
                queue = append(queue, &Position{row: row, col: col})
            }
        }
    }

    if fresh == 0 {
        return 0
    }

    mins := -1

    for len(queue) > 0 {
        next := []*Position{}

        for _, pos := range queue {
            if target := pos.Up(); rot(grid, target) {
                next = append(next, target)
                fresh--
            }

            if target := pos.Down(); rot(grid, target) {
                next = append(next, target)
                fresh--
            }

            if target := pos.Left(); rot(grid, target) {
                next = append(next, target)
                fresh--
            }

            if target := pos.Right(); rot(grid, target) {
                next = append(next, target)
                fresh--
            }
        }

        queue = next
        mins++
    }

    if fresh > 0 {
        return -1
    }

    return mins
}
