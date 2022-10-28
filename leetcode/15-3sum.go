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

    for num1 := range bucket {
        bucket[num1]--

        for num2 := range bucket {
            target := 0 - num1 - num2

            if _, ok := bucket[target]; !ok {
                continue
            } else if bucket[num2] < 1 {
                continue
            } else if bucket[target] < 1 {
                continue
            } else if num2 == target && bucket[num2] < 2 {
                continue
            }

            seen[makeTriplet(num1, num2, target)] = true
        }

        bucket[num1]++
    }

    out := make([][]int, 0, len(seen))

    for triplet := range seen {
        triplet := []int{triplet[0], triplet[1], triplet[2]}

        out = append(out, triplet)
    }

    return out
}
