func findMedianSorted(list []int, size int) float64 {
    mid := size / 2
    
    switch {
    case size == 0:
        return 0
    case size & 0x1 == 1:
        return float64(list[mid])
    default:
        return float64(list[mid] + list[mid-1]) / 2
    }
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) == 0 {
        return findMedianSorted(nums2, len(nums2))
    } else if len(nums2) == 0 {
        return findMedianSorted(nums1, len(nums1))
    }

    size := len(nums1) + len(nums2)
    capapicy := size / 2 + 1
    merged := make([]int, 0, capapicy)

    for len(merged) < capapicy {
        num := 0

        switch {
        case len(nums1) == 0:
            num, nums2 = nums2[0], nums2[1:]
        case len(nums2) == 0:
            num, nums1 = nums1[0], nums1[1:]
        case nums1[0] < nums2[0]:
            num, nums1 = nums1[0], nums1[1:]
        default:
            num, nums2 = nums2[0], nums2[1:]
        }

        merged = append(merged, num)
    }

    return findMedianSorted(merged, size)
}
