// #array

func distinct(nums []int, ref []int) []int {
    out := make([]int, 0, len(nums))
    seen := map[int]bool{}

    for _, num := range ref {
        seen[num] = true
    }

    for _, num := range nums {
        if _, have := seen[num]; !have {
            seen[num] = true
            out = append(out, num)
        }
    }

    return out
}

func findDifference(nums1 []int, nums2 []int) [][]int {
    out := [][]int{
        distinct(nums1, nums2),
        distinct(nums2, nums1),
    }

    return out
}
