func findKthPositive(arr []int, k int) int {
    left := 0
    right := len(arr)-1

    for left <= right {
        mid := left + (right - left)/2

        if arr[mid] - (mid+1) < k {
            left = mid+1
        } else {
            right = mid-1
        }
    }

    if right == -1 {
        return k
    } 

    return k + right + 1
}