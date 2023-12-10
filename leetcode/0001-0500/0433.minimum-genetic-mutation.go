// #graph #shorest-path

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func mutation(lhs string, rhs string) int {
    out := 0

    for i := range lhs {
        if lhs[i] != rhs[i] {
            out++
        }
    }

    if out > 1 {
        return -1
    }

    return out
}

func makeGraph(genes []string) [][]int {
    out := make([][]int, len(genes))

    for i := range out {
        out[i] = make([]int, len(genes))
    }

    for from := 0; from < len(out); from++ {
        for to := from; to < len(out); to++ {
            out[from][to] = mutation(genes[from], genes[to])
            out[to][from] = out[from][to]
        }
    }

    return out
}

func minMutation(start string, end string, bank []string) int {
    if len(bank) == 0 {
        return -1
    }

    bank = append(bank, start)
    distance := makeGraph(bank)

    for mid := 0; mid < len(distance); mid++ {
        for from := 0; from < len(distance); from++ {
            for to := 0; to < len(distance); to++ {
                if distance[from][mid] == -1 {
                    continue
                } else if distance[mid][to] == -1 {
                    continue
                }

                dist := distance[from][mid] + distance[mid][to]

                if distance[from][to] == -1 || distance[from][to] > dist {
                    distance[from][to] = dist
                }
            }
        }
    }

    for to, gene := range bank {
        if gene == end {
            return distance[len(bank)-1][to]
        }
    }

    return -1
}
