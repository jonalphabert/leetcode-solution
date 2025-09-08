func maximalSquare(matrix [][]byte) int {
    n, m := len(matrix), len(matrix[0])
    dp := make([][]int, n+1)
    maxSide := 0

    for i:=0; i<=n; i++ {
        dp[i] = make([]int, m+1)
    }

    for i:=1; i<=n; i++ {
        for j:=1; j<=m; j++ {
            if matrix[i-1][j-1] == '1' {
                dp[i][j] = min(dp[i][j-1], min(dp[i-1][j], dp[i-1][j-1])) + 1
                maxSide = max(maxSide, dp[i][j])
            }
        }
    }

    return maxSide * maxSide
}