/* Information for this solution:
 * This solution inspired by BFS algorithm with DP array to store the minimum jumps to reach each index.
 * Therefore, there is another better solution with O(n) time complexity.
 * Complexity: O(n^2)
 * Space: O(n)
*/
import "math"

func jump(nums []int) int {
    length := len(nums)
    dp := make([] int, length)

    // Iniasisi dp
    for i := range dp {
        dp[i] = math.MaxInt
    }

    dp[0] = 0
    isReachLastIndex := false
    for i := 0; i < length; i++ {
        for j := 1; j <= nums[i] && i+j < length; j++ {
            dp[i+j] = min(dp[i+j], dp[i]+1)

            if i+j == length - 1 {
                isReachLastIndex = true
                break
            }
        }
        if isReachLastIndex == true {
            break
        }
    }

    return dp[length-1]
}