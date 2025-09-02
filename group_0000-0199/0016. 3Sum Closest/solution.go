import (
    "math"
    "sort"
)

func threeSumClosest(nums []int, target int) int {
    n := len(nums)
    minDiff := math.MaxInt
    result := math.MaxInt

    sort.Slice(nums, func (l,r int) bool {
        return nums[l] < nums[r]
    })

    for i := 0; i < n; i++ {
        left := i+1
        right := n-1

        for left < right {
            currentResult := nums[i] + nums[left] + nums[right]

            if minDiff > int(math.Abs(float64(currentResult - target))) {
                minDiff = int(math.Abs(float64(currentResult - target)))
                result = currentResult
            }

            if currentResult > target {
                right--
            } else {
                left++
            }
        }
    }

    return result
}