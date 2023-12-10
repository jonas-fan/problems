// #array

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    out := []int{}

    for _, num := range nums1 {
        index := 0

        for ; index < len(nums2); index++ {
            if nums2[index] == num {
                break
            }
        }

        for ; index < len(nums2); index++ {
            if nums2[index] > num {
                break
            }
        }

        if index == len(nums2) {
            out = append(out, -1)
        } else {
            out = append(out, nums2[index])
        }
    }

    return out
}
