package main

func merge1(nums1 []int, m int, nums2 []int, n int) {
    index := len(nums1) - 1
    for m >= 1 && n >= 1 {
        if nums1[m-1] > nums2[n-1] {
            nums1[index] = nums1[m-1]
            m--
        } else {
            nums1[index] = nums2[n-1]
            n--
        }
        index--
    }
    for m >= 1 {
        nums1[index] = nums1[m-1]
        m--
        index--
    }
    for n >= 1 {
        nums1[index] = nums2[n-1]
        n--
        index--
    }
}
