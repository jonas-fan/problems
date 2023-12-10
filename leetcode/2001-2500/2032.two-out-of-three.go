// #array

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    seen1 := [101]bool{}
    seen2 := [101]bool{}
    seen3 := [101]bool{}

    for _, num := range nums1 {
        seen1[num] = true
    }

    for _, num := range nums2 {
        seen2[num] = true
    }

    for _, num := range nums3 {
        seen3[num] = true
    }

    out := []int{}

    for i := 1; i < 101; i++ {
        if seen1[i] && seen2[i] {
            out = append(out, i)
        } else if seen1[i] && seen3[i] {
            out = append(out, i)
        } else if seen2[i] && seen3[i] {
            out = append(out, i)
        }
    }

    return out
}
