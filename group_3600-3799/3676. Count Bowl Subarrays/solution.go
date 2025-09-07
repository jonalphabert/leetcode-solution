func bowlSubarrays(nums []int) int64 {
    stack := []int{}
    result := int64(0)

    for i := 0; i < len(nums); i++ {
        for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
            stack = stack[:len(stack) - 1]

            if len(stack) == 0 {
                break
            }
            result++
        }

        stack = append(stack, i)
    }

    return result
}