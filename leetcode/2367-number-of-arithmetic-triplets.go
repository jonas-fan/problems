// #array

func arithmeticTriplets(nums []int, diff int) int {
    out := 0
    seen := map[int]int{}

    for i := 0; i < len(nums); i++ {
        seen[nums[i]] = i
    }

    for i := 0; i < len(nums); i++ {
        if _, ok := seen[nums[i]+diff]; !ok {
            continue
        } else if _, ok := seen[nums[i]+diff+diff]; !ok {
            continue
        }

        out++
    }

    return out
}
