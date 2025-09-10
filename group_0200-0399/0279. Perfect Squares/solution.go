func numSquares(n int) int {
    dp := make([]int, n+1)
    var minimumBuilt int

    dp[1] = 1

    for i := 2; i<=n; i++ {
        minimumBuilt = math.MaxInt
        for j := 1; j * j <= i; j++ {
            minimumBuilt = min(minimumBuilt, dp[i-j*j] + 1)
        }
        dp[i] = minimumBuilt
    }

    return dp[n]
}