// #array

import (
    "sort"
)

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)

    out := make([]int, 0, max(len(nums1), len(nums2)))

    for len(nums1) > 0 && len(nums2) > 0 {
        if nums1[0] < nums2[0] {
            nums1 = nums1[1:]
        } else if nums2[0] < nums1[0] {
            nums2 = nums2[1:]
        } else if nums1[0] == nums2[0] {
            out = append(out, nums1[0])
            nums1 = nums1[1:]
            nums2 = nums2[1:]
        }
    }

    return out
}
