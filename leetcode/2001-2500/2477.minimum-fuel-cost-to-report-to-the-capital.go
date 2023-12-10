// #graph #depth-first-search

func dfs(graph map[int][]int, city int, last int, seats int) (int, int) {
    count := 1
    fuel := 0

    for _, next := range graph[city] {
        if next != last {
            nextCount, nextFuel := dfs(graph, next, city, seats)

            count += nextCount
            fuel += nextFuel
        }
    }

    if city > 0 {
        fuel += (count + seats - 1) / seats
    }

    return count, fuel
}

func minimumFuelCost(roads [][]int, seats int) int64 {
    graph := map[int][]int{}

    for _, road := range roads {
        graph[road[0]] = append(graph[road[0]], road[1])
        graph[road[1]] = append(graph[road[1]], road[0])
    }

    _, fuel := dfs(graph, 0, -1, seats)

    return int64(fuel)
}
