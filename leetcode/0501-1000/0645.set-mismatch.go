func findErrorNums(nums []int) []int {
    rest := (1 + len(nums)) * len(nums) / 2
    seen := map[int]int{}
    dup := 0

    for _, num := range nums {
        if _, ok := seen[num]; ok {
            dup = num
            continue
        }

        seen[num]++
        rest -= num
    }

    return []int{dup, rest}
}
