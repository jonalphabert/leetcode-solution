func nextGreaterElement(nums1 []int, nums2 []int) []int {
    nextGreater := make(map[int]int)
    stack := []int{}

    for i := 0; i < len(nums2); i++ {
        for len(stack) > 0 && stack[len(stack) - 1] < nums2[i] {
            // Pop stack
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            // Make a map from top variable with nums[i]
            nextGreater[top] = nums2[i]
        }

        stack = append(stack, nums2[i])
    }

    res := []int{}

    for i := 0; i < len(nums1); i++ {
        if val, exist := nextGreater[nums1[i]]; exist {
            res = append(res, val)
        } else {
            res = append(res, -1)
        }
    }

    return res
}