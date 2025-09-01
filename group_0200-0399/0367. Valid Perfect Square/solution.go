func isPerfectSquare(num int) bool {
    left := 1
    right := num

    for left <= right {
        mid := left + (right-left)/2

        power2mid := mid * mid

        if power2mid == num {
            return true
        }

        if num > power2mid {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return false
}