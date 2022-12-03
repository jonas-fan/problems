// #array

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func maxDistance(nums1 []int, nums2 []int) int {
    out := 0
    i := 0
    j := 0

    for i < len(nums1) && j < len(nums2) {
        if nums1[i] <= nums2[j] {
            out = max(out, j-i)
            j++
        } else {
            i++
        }
    }

    return out
}
