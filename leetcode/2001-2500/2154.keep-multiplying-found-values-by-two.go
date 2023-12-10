// #array

func findFinalValue(nums []int, original int) int {
    seen := map[int]int{}

    for _, num := range nums {
        seen[num]++
    }

    for {
        if _, have := seen[original]; !have {
            break
        }

        original <<= 1
    }

    return original
}
