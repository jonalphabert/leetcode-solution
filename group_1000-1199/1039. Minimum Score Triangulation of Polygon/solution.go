func minScoreTriangulation(values []int) int {
    n := len(values)
    dp := make([][]int, n)

    for u:=0; u<n; u++ {
        dp[u] = make([]int, n)
    }

    for i := n-1; i >= 0; i-- {
        for j := i+2; j < n; j++ {
            dp[i][j] = 1<<31 - 1
            for k := i+1; k < j; k++ {
                dp[i][j] = min(dp[i][j], dp[i][k] + dp[k][j] + values[i] * values[j] * values[k])
            }
        }
    }

    return dp[0][n-1]
}