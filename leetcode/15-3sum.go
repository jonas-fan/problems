import (
    "sort"
)

func makeTriplet(x int, y int, z int) [3]int {
    out := [3]int{x, y, z}

    sort.Ints(out[:])

    return out
}

func threeSum(nums []int) [][]int {
    bucket := map[int]int{}

    for _, num := range nums {
        bucket[num]++
    }

    seen := map[[3]int]bool{}

    for x := range bucket {
        bucket[x]--

        for y := range bucket {
            z := 0 - x - y

            if _, ok := bucket[z]; !ok {
                continue
            } else if bucket[y] < 1 {
                continue
            } else if bucket[z] < 1 {
                continue
            } else if y == z && bucket[y] < 2 {
                continue
            }

            seen[makeTriplet(x, y, z)] = true
        }

        bucket[x]++
    }

    out := make([][]int, 0, len(seen))

    for triplet := range seen {
        triplet := []int{triplet[0], triplet[1], triplet[2]}

        out = append(out, triplet)
    }

    return out
}
