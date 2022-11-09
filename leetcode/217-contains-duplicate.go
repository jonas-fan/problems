// #array #hashmap

func containsDuplicate(nums []int) bool {
    seen := map[int]bool{}

    for _, num := range nums {
        if _, have := seen[num]; have {
            return true
        }

        seen[num] = true
    }

    return false
}
