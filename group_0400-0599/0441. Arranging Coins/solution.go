func arrangeCoins(n int) int {
    left := 1
    right := n

    for left <= right {
        mid := left + (right - left) /2

        staircase := mid * (mid+1) / 2

        if n == staircase {
            return mid
        }

        if n < staircase {
            right = mid -1
        } else {
            left = mid + 1
        }
    }

    return right
}