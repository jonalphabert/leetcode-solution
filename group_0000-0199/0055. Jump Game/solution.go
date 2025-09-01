func canJump(nums []int) bool {
    n := len(nums)
    farthest := 0
    end := 0

    for i := 0; i < n-1 && i <= end; i++ {
        farthest = max(farthest, i+nums[i])
        if i == end {
            end = farthest
        }
    }

    return farthest >= n-1
}