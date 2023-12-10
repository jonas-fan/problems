// #array

func merge(nums1 []int, m int, nums2 []int, n int)  {
    end := m + n

    for m > 0 || n > 0 {
        num := 0

        if m == 0 {
            n--
            num = nums2[n]
        } else if n == 0 {
            m--
            num = nums1[m]
        } else if nums1[m-1] < nums2[n-1] {
            n--
            num = nums2[n]
        } else {
            m--
            num = nums1[m]
        }

        end--
        nums1[end] = num
    }
}
