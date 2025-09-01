import (
    "math"
)

func findMin(nums []int) int {
    size := len(nums)
    left := 0
    right := size-1
    mid := 0
    m := math.MaxInt

    // Mencari nilai terkecil dari array
    for left <= right {
        mid = left + (right - left)/2

        if nums[left] == nums[right] {
            m = min(m, nums[left])
            left++
            right--
            continue
        }

        // Jika angka di mid masih lebih besar daripada angka di left
        // maka bisa dipastikan semua nilai di sebelah kiri mid sudah terurut
        // Jika kebalikannya, maka partisi dari array yang terrotasi ada di antara left dan juga mid
        if nums[left] <= nums[mid] {
            m = min(m, nums[left])
            left = mid + 1
        } else {
            m = min(m, nums[mid])
            right = mid
        }
    }

    return m
}