func maxTotalValue(nums []int, k int) int64 {
    maxIndex := 0
    minIndex := 0

    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[maxIndex] {
            maxIndex = i
        }
        if nums[i] < nums[minIndex] {
            minIndex = i
        }
    }

    return int64(k) * int64(nums[maxIndex] - nums[minIndex])
}