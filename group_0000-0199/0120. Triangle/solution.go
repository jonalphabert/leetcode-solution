func minimumTotal(triangle [][]int) int {
    minValue := math.MaxInt

    for i := 1; i < len(triangle) ; i++ {
        for j := 0; j < len(triangle[i]); j++ {
            nums1 := math.MaxInt
            nums2 := math.MaxInt
            if j != 0 {
                nums1 = triangle[i-1][j-1]
            }
            if j != len(triangle[i])-1 {
                nums2 = triangle[i-1][j]
            }

            triangle[i][j] += min(nums1, nums2)
        }
    }

    for _, val := range triangle[len(triangle) -1] {
        minValue = min(minValue, val)
    }

    return minValue
}